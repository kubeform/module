/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"bytes"
	"compress/gzip"
	"context"
	"ekyu.moe/base91"
	"encoding/json"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/fatih/structs"
	"github.com/gobuffalo/flect"
	"gocloud.dev/secrets"
	"io/ioutil"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	kmapi "kmodules.xyz/client-go/api/v1"
	"kmodules.xyz/client-go/meta"
	"os"
	"path/filepath"
	"reflect"
	"sigs.k8s.io/cli-utils/pkg/kstatus/status"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"strings"
)

var SecretKey string

func StartProcess(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	err := initialUpdateStatus(rClient, ctx, gv, obj, nil, true)
	if err != nil {
		fmt.Println("1")
		return err
	}

	err = reconcile(rClient, ctx, gv, obj)
	if err != nil {
		err2 := initialUpdateStatus(rClient, ctx, gv, obj, err, false)
		if err2 != nil {
			fmt.Println("2")
			return err2
		}
		fmt.Println("3")
		return err
	}

	return finalUpdateStatus(rClient, ctx, obj)
}

func reconcile(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	providerName, found, err := unstructured.NestedString(obj.Object, "spec", "providerName")
	if err != nil {
		return err
	}
	if !found {
		return fmt.Errorf("providerName is not found")
	}

	// let's suppose moduleDef is the module source for now, in future would do it from ModuleDefinition
	moduleDef, found, err := unstructured.NestedString(obj.Object, "spec", "moduleDef", "name")
	if err != nil {
		fmt.Println("4")
		return err
	}
	if !found {
		return fmt.Errorf("moduleDef is not found")
	}

	// TODO: validation part
	// when we will have module def created then we can use this
	// demo of validation is done before in the tf-module-support-in-kf repo
	// now as we supposed moduleDef is the module source for now, so just ignore this validation part for now

	namespace := obj.GetNamespace()
	moduleName := obj.GetName()
	resPath := filepath.Join(basePath, "modules"+"."+namespace+"."+moduleName)
	mainFile := filepath.Join(resPath, "main.tf.json")
	stateFile := filepath.Join(resPath, "terraform.tfstate")
	outputFile := filepath.Join(resPath, "output.tf")

	if hasFinalizer(obj.GetFinalizers(), KFCFinalizer) {
		if obj.GetDeletionTimestamp() != nil {
			err := updateStatus(rClient, ctx, obj, status.TerminatingStatus)
			if err != nil {
				return err
			}

			err = terraformDestroy(resPath)
			if err != nil {
				return err
			}

			err = deleteFiles(resPath)
			if err != nil {
				return err
			}

			return removeFinalizer(ctx, rClient, obj, KFCFinalizer)
		}
	} else {
		err := addFinalizer(ctx, rClient, obj, KFCFinalizer)
		if err != nil {
			return err
		}
	}
	fmt.Println("oka let's start")
	err = createFiles(resPath, mainFile)
	if err != nil {
		return err
	}
	fmt.Println("before mainTFJson")
	mainTfJson, err := mainTFJson(rClient, ctx, moduleDef, providerName, moduleName, obj)
	err = ioutil.WriteFile(mainFile, mainTfJson, 0777)
	if err != nil {
		return err
	}
	fmt.Println("before terraformInit")
	err = terraformInit(resPath)
	if err != nil {
		return err
	}
	fmt.Println("before createTFStteFile")
	err = createTFStateFile(stateFile, gv, obj)
	if err != nil {
		return err
	}
	fmt.Println("before terraformApply")
	err = terraformApply(resPath)
	if err != nil {
		return err
	}
	fmt.Println("before updateTFStateFile")
	err = updateTFStateFile(stateFile, gv, obj)
	if err != nil {
		return err
	}
	fmt.Println("before updateOutputField")
	err = updateOutputField(resPath, moduleName, outputFile, gv, obj)
	if err != nil {
		return err
	}

	return nil
}

func mainTFJson(rClient client.Client, ctx context.Context, source, providerName, moduleName string, obj *unstructured.Unstructured) ([]byte, error) {
	spew.Dump(obj.Object)
	input, found, err := unstructured.NestedMap(obj.Object, "spec", "resource", "input")
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("no input is found")
	}

	pureInput := make(map[string]interface{})
	fmt.Println("let's see the input")
	spew.Dump(input)
	fmt.Println("after printing input")
	for key, val := range input {
		pureInput[flect.Underscore(key)] = val
	}
	pureInput["source"] = source
	fmt.Println("let's see the pure input")
	spew.Dump(pureInput)
	fmt.Println("after prinint pure input")
	jsnInput, err := json.Marshal(pureInput)
	if err != nil {
		return nil, err
	}

	finalJson := []byte(`{`)

	// now hardcoded providerSource, later will read from ModuleDef
	finalJson = append(finalJson, []byte(`"terraform": {
		"required_providers": {
			"`+providerName+`": {
				"source": "terraform-providers/linode"
			}
		}
	},`)...)

	providerRef, found, err := unstructured.NestedString(obj.Object, "spec", "providerRef", "name")
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, fmt.Errorf("providerRef is not found")
	}
	var secret corev1.Secret
	request := types.NamespacedName{
		Namespace: obj.GetNamespace(),
		Name:      providerRef,
	}
	if err := rClient.Get(ctx, request, &secret); err != nil {
		return nil, err
	}

	providerSecretData := secret.Data["provider"]

	finalJson = append(finalJson, []byte(`"provider": { "`+providerName+`": `)...)
	finalJson = append(finalJson, providerSecretData...)
	finalJson = append(finalJson, []byte(`},`)...)

	moduleData := []byte(`{"` + moduleName + `":`)
	moduleData = append(moduleData, jsnInput...)
	moduleData = append(moduleData, []byte("}")...)
	prettyData, err := prettyJSON(moduleData)
	if err != nil {
		return nil, err
	}

	finalJson = append(finalJson, []byte(`"module": `)...)
	finalJson = append(finalJson, prettyData...)
	finalJson = append(finalJson, []byte(`}`)...)

	return finalJson, nil
}

func deleteFiles(resPath string) error {
	err := os.RemoveAll(resPath)
	if err != nil {
		return err
	}

	return nil
}

func initialUpdateStatus(rClient client.Client, ctx context.Context, gv schema.GroupVersion, obj *unstructured.Unstructured, er error, flag bool) error {
	objGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
	if err != nil {
		return err
	}

	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}

	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	conditionsVal := reflect.ValueOf(typedStruct.Field("Status").Field("Conditions").Value())
	conditions := conditionsVal.Interface().([]kmapi.Condition)
	if kmapi.HasCondition(conditions, "Stalled") {
		return nil
	}

	phase := status.InProgressStatus
	if flag {
		conditions = kmapi.SetCondition(conditions, kmapi.NewCondition("Reconciling", "Kubeform is currently reconciling "+obj.GetKind()+" resource", objGen))
	} else {
		conditions = kmapi.SetCondition(conditions, kmapi.NewCondition("Stalled", er.Error(), objGen))
		phase = status.FailedStatus
	}

	err = setNestedFieldNoCopy(obj.Object, conditions, "status", "conditions")
	if err != nil {
		return err
	}
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}

	return updateStatus(rClient, ctx, obj, phase)
}

func finalUpdateStatus(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured) error {
	var newCondi []kmapi.Condition
	err := setNestedFieldNoCopy(obj.Object, newCondi, "status", "conditions")
	if err != nil {
		return err
	}
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	err = updateStatus(rClient, ctx, obj, status.CurrentStatus)
	if err != nil {
		return err
	}
	return nil
}

func updateStatus(rClient client.Client, ctx context.Context, obj *unstructured.Unstructured, phase status.Status) error {
	if phase == status.CurrentStatus {
		obsGen, _, err := unstructured.NestedInt64(obj.Object, "metadata", "generation")
		if err != nil {
			return err
		}
		err = unstructured.SetNestedField(obj.Object, obsGen, "status", "observedGeneration")
		if err != nil {
			return err
		}
	}

	err := setNestedFieldNoCopy(obj.Object, phase, "status", "phase")
	if err != nil {
		return err
	}

	// apply the status update of the object
	if err = rClient.Status().Update(ctx, obj); err != nil {
		return err
	}
	return nil
}

func setNestedFieldNoCopy(obj map[string]interface{}, value interface{}, fields ...string) error {
	m := obj

	for i, field := range fields[:len(fields)-1] {
		if val, ok := m[field]; ok {
			if valMap, ok := val.(map[string]interface{}); ok {
				m = valMap
			} else {
				return fmt.Errorf("value cannot be set because %v is not a map[string]interface{}", jsonPath(fields[:i+1]))
			}
		} else {
			newVal := make(map[string]interface{})
			m[field] = newVal
			m = newVal
		}
	}
	m[fields[len(fields)-1]] = value
	return nil
}

func jsonPath(fields []string) string {
	return "." + strings.Join(fields, ".")
}

func prettyJSON(byteJson []byte) ([]byte, error) {
	var prettyJSON bytes.Buffer
	err := json.Indent(&prettyJSON, byteJson, "", "  ")
	if err != nil {
		return nil, err
	}

	return prettyJSON.Bytes(), err
}

func hasFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}

	return false
}

func addFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for _, v := range finalizers {
		if v == finalizer {
			return nil
		}
	}
	finalizers = append(finalizers, finalizer)
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}
	err = rClient.Update(ctx, u)
	return err
}

func removeFinalizer(ctx context.Context, rClient client.Client, u *unstructured.Unstructured, finalizer string) error {
	finalizers := u.GetFinalizers()
	for i, v := range finalizers {
		if v == finalizer {
			finalizers = append(finalizers[:i], finalizers[i+1:]...)
			break
		}
	}
	err := unstructured.SetNestedStringSlice(u.Object, finalizers, "metadata", "finalizers")
	if err != nil {
		return err
	}

	err = rClient.Update(ctx, u)
	return err
}

func createFiles(resPath, mainFile string) error {
	_, err := os.Stat(resPath)
	if os.IsNotExist(err) {
		err := os.MkdirAll(resPath, 0777)
		if err != nil {
			return err
		}

		_, err = os.Create(mainFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func createTFStateFile(filePath string, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	data, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}
	typedObj, err := meta.UnmarshalFromJSON(data, gv)
	if err != nil {
		return err
	}
	typedStruct := structs.New(typedObj)
	stateValue := typedStruct.Field("Spec").Field("State").Value()

	_, existErr := os.Stat(filePath)
	if os.IsNotExist(existErr) && stateValue.(string) != "" {
		decodedData, err := decodeState(stateValue.(string))
		if err != nil {
			return err
		}

		err = ioutil.WriteFile(filePath, decodedData, 0644)
		if err != nil {
			return fmt.Errorf("failed to write file hash : %s", err.Error())
		}
	}

	return nil
}

func updateTFStateFile(filePath string, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	jsonData, err := meta.MarshalToJson(obj, gv)
	if err != nil {
		return err
	}
	typedObj, err := meta.UnmarshalFromJSON(jsonData, gv)
	if err != nil {
		return err
	}

	typedStruct := structs.New(typedObj)
	stateValue := typedStruct.Field("Spec").Field("State").Value()

	if stateValue.(string) == "" || !reflect.DeepEqual([]byte(stateValue.(string)), data) {
		processedData, err := encodeState(data)
		if err != nil {
			return err
		}

		err = unstructured.SetNestedField(obj.Object, processedData, "spec", "state")
		if err != nil {
			return fmt.Errorf("failed to update spec state : %s", err)
		}

	}
	return nil
}

func decodeState(data string) ([]byte, error) {
	cipherText := base91.DecodeString(data)

	savedKeyKeeper, err := secrets.OpenKeeper(context.Background(), "base64key://"+SecretKey)
	if err != nil {
		return nil, err
	}
	defer savedKeyKeeper.Close()

	plainText, err := savedKeyKeeper.Decrypt(context.Background(), cipherText)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(plainText)

	zr, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(zr)
	if err != nil {
		return nil, err
	}

	if err := zr.Close(); err != nil {
		return nil, err
	}

	return result, nil
}

func encodeState(data []byte) (string, error) {
	// zip
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	if _, err := zw.Write(data); err != nil {
		return "", err
	}

	if err := zw.Close(); err != nil {
		return "", err
	}

	// encrypt
	savedKeyKeeper, err := secrets.OpenKeeper(context.Background(), "base64key://"+SecretKey)
	if err != nil {
		return "", err
	}
	defer savedKeyKeeper.Close()

	cipherText, err := savedKeyKeeper.Encrypt(context.Background(), buf.Bytes())
	if err != nil {
		return "", err
	}

	// base91

	return base91.EncodeToString(cipherText), nil
}

func updateOutputField(resPath, outputFile, moduleName string, gv schema.GroupVersion, obj *unstructured.Unstructured) error {
	_, err := os.Stat(outputFile)
	if os.IsNotExist(err) {
		data, err := meta.MarshalToJson(obj, gv)
		if err != nil {
			return err
		}

		typedObj, err := meta.UnmarshalFromJSON(data, gv)
		if err != nil {
			return err
		}

		typedStruct := structs.New(typedObj)
		outputData := []byte(``)
		output := reflect.TypeOf(typedStruct.Field("Spec").Field("Output").Value()).Elem()

		for i := 0; i < output.NumField(); i++ {
			field := output.Field(i).Tag.Get("tf")
			outputData = append(outputData, []byte(`output "`+field+`" { 
value = module.`+moduleName+`.`+field+` 
}
`)...)
		}

		err = ioutil.WriteFile(outputFile, outputData, 0644)
		if err != nil {
			return err
		}
	}

	value, err := terraformOutput(resPath)
	if err != nil {
		return err
	}

	outputs := make(map[string]output)

	err = json.Unmarshal([]byte(value), &outputs)
	if err != nil {
		return err
	}

	for name, output := range outputs {
		val, err := output.ValueRaw.MarshalJSON()
		if err != nil {
			return err
		}

		err = setNestedFieldNoCopy(obj.Object, string(val), "spec", "resource", "output", flect.Camelize(name))
		if err != nil {
			return err
		}
	}

	return nil
}
