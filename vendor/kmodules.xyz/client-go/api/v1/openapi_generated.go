//go:build !ignore_autogenerated
// +build !ignore_autogenerated

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

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"kmodules.xyz/client-go/api/v1.CertificatePrivateKey": schema_kmodulesxyz_client_go_api_v1_CertificatePrivateKey(ref),
		"kmodules.xyz/client-go/api/v1.CertificateSpec":       schema_kmodulesxyz_client_go_api_v1_CertificateSpec(ref),
		"kmodules.xyz/client-go/api/v1.ClusterMetadata":       schema_kmodulesxyz_client_go_api_v1_ClusterMetadata(ref),
		"kmodules.xyz/client-go/api/v1.Condition":             schema_kmodulesxyz_client_go_api_v1_Condition(ref),
		"kmodules.xyz/client-go/api/v1.ObjectID":              schema_kmodulesxyz_client_go_api_v1_ObjectID(ref),
		"kmodules.xyz/client-go/api/v1.ObjectInfo":            schema_kmodulesxyz_client_go_api_v1_ObjectInfo(ref),
		"kmodules.xyz/client-go/api/v1.ObjectReference":       schema_kmodulesxyz_client_go_api_v1_ObjectReference(ref),
		"kmodules.xyz/client-go/api/v1.ResourceID":            schema_kmodulesxyz_client_go_api_v1_ResourceID(ref),
		"kmodules.xyz/client-go/api/v1.TLSConfig":             schema_kmodulesxyz_client_go_api_v1_TLSConfig(ref),
		"kmodules.xyz/client-go/api/v1.TimeOfDay":             schema_kmodulesxyz_client_go_api_v1_TimeOfDay(ref),
		"kmodules.xyz/client-go/api/v1.TypedObjectReference":  schema_kmodulesxyz_client_go_api_v1_TypedObjectReference(ref),
		"kmodules.xyz/client-go/api/v1.X509Subject":           schema_kmodulesxyz_client_go_api_v1_X509Subject(ref),
		"kmodules.xyz/client-go/api/v1.stringSetMerger":       schema_kmodulesxyz_client_go_api_v1_stringSetMerger(ref),
	}
}

func schema_kmodulesxyz_client_go_api_v1_CertificatePrivateKey(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CertificatePrivateKey contains configuration options for private keys used by the Certificate controller. This allows control of how private keys are rotated.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"encoding": {
						SchemaProps: spec.SchemaProps{
							Description: "The private key cryptography standards (PKCS) encoding for this certificate's private key to be encoded in. If provided, allowed values are \"pkcs1\" and \"pkcs8\" standing for PKCS#1 and PKCS#8, respectively. Defaults to PKCS#1 if not specified. See here for the difference between the formats: https://stackoverflow.com/a/48960291",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_CertificateSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"alias": {
						SchemaProps: spec.SchemaProps{
							Description: "Alias represents the identifier of the certificate.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"issuerRef": {
						SchemaProps: spec.SchemaProps{
							Description: "IssuerRef is a reference to a Certificate Issuer.",
							Ref:         ref("k8s.io/api/core/v1.TypedLocalObjectReference"),
						},
					},
					"secretName": {
						SchemaProps: spec.SchemaProps{
							Description: "Specifies the k8s secret name that holds the certificates. Default to <resource-name>-<cert-alias>-cert.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"subject": {
						SchemaProps: spec.SchemaProps{
							Description: "Full X509 name specification (https://golang.org/pkg/crypto/x509/pkix/#Name).",
							Ref:         ref("kmodules.xyz/client-go/api/v1.X509Subject"),
						},
					},
					"duration": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate default Duration",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
					"renewBefore": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate renew before expiration duration",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Duration"),
						},
					},
					"dnsNames": {
						SchemaProps: spec.SchemaProps{
							Description: "DNSNames is a list of subject alt names to be used on the Certificate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"ipAddresses": {
						SchemaProps: spec.SchemaProps{
							Description: "IPAddresses is a list of IP addresses to be used on the Certificate",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"uris": {
						SchemaProps: spec.SchemaProps{
							Description: "URIs is a list of URI subjectAltNames to be set on the Certificate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"emailAddresses": {
						SchemaProps: spec.SchemaProps{
							Description: "EmailAddresses is a list of email subjectAltNames to be set on the Certificate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"privateKey": {
						SchemaProps: spec.SchemaProps{
							Description: "Options to control private keys used for the Certificate.",
							Ref:         ref("kmodules.xyz/client-go/api/v1.CertificatePrivateKey"),
						},
					},
				},
				Required: []string{"alias"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.TypedLocalObjectReference", "k8s.io/apimachinery/pkg/apis/meta/v1.Duration", "kmodules.xyz/client-go/api/v1.CertificatePrivateKey", "kmodules.xyz/client-go/api/v1.X509Subject"},
	}
}

func schema_kmodulesxyz_client_go_api_v1_ClusterMetadata(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"uid": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"displayName": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"provider": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
				Required: []string{"uid"},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_Condition(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of condition in CamelCase or in foo.example.com/CamelCase. Many .condition.type values are consistent across resources like Available, but because arbitrary conditions can be useful (see .node.status.conditions), the ability to deconflict is important.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"observedGeneration": {
						SchemaProps: spec.SchemaProps{
							Description: "If set, this represents the .metadata.generation that the condition was set based upon. For instance, if .metadata.generation is currently 12, but the .status.condition[x].observedGeneration is 9, the condition is out of date with respect to the current state of the instance.",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
					"lastTransitionTime": {
						SchemaProps: spec.SchemaProps{
							Description: "Last time the condition transitioned from one status to another. This should be when the underlying condition changed.  If that is not known, then using the time when the API field changed is acceptable.",
							Default:     map[string]interface{}{},
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"reason": {
						SchemaProps: spec.SchemaProps{
							Description: "The reason for the condition's last transition in CamelCase. The specific API may choose whether or not this field is considered a guaranteed API. This field may not be empty.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition. This field may be empty.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"type", "status", "lastTransitionTime", "reason", "message"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_kmodulesxyz_client_go_api_v1_ObjectID(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_ObjectInfo(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"resource": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("kmodules.xyz/client-go/api/v1.ResourceID"),
						},
					},
					"ref": {
						SchemaProps: spec.SchemaProps{
							Default: map[string]interface{}{},
							Ref:     ref("kmodules.xyz/client-go/api/v1.ObjectReference"),
						},
					},
				},
				Required: []string{"resource", "ref"},
			},
		},
		Dependencies: []string{
			"kmodules.xyz/client-go/api/v1.ObjectReference", "kmodules.xyz/client-go/api/v1.ResourceID"},
	}
}

func schema_kmodulesxyz_client_go_api_v1_ObjectReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ObjectReference contains enough information to let you inspect or modify the referred object.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_ResourceID(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceID identifies a resource",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"group": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name is the plural name of the resource to serve.  It must match the name of the CustomResourceDefinition-registration too: plural.group and it must be all lowercase.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is the serialized kind of the resource.  It is normally CamelCase and singular.",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"scope": {
						SchemaProps: spec.SchemaProps{
							Default: "",
							Type:    []string{"string"},
							Format:  "",
						},
					},
				},
				Required: []string{"group", "version", "name", "kind", "scope"},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_TLSConfig(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"issuerRef": {
						SchemaProps: spec.SchemaProps{
							Description: "IssuerRef is a reference to a Certificate Issuer.",
							Ref:         ref("k8s.io/api/core/v1.TypedLocalObjectReference"),
						},
					},
					"certificates": {
						SchemaProps: spec.SchemaProps{
							Description: "Certificate provides server and/or client certificate options used by application pods. These options are passed to a cert-manager Certificate object. xref: https://github.com/jetstack/cert-manager/blob/v0.16.0/pkg/apis/certmanager/v1beta1/types_certificate.go#L82-L162",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: map[string]interface{}{},
										Ref:     ref("kmodules.xyz/client-go/api/v1.CertificateSpec"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.TypedLocalObjectReference", "kmodules.xyz/client-go/api/v1.CertificateSpec"},
	}
}

func schema_kmodulesxyz_client_go_api_v1_TimeOfDay(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TimeOfDay is a wrapper around time.Time which supports correct marshaling to YAML and JSON.  Wrappers are provided for many of the factory methods that the time package offers.",
				Type:        TimeOfDay{}.OpenAPISchemaType(),
				Format:      TimeOfDay{}.OpenAPISchemaFormat(),
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_TypedObjectReference(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "TypedObjectReference represents an typed namespaced object.",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"apiGroup": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"kind": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name of the referent. More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names",
							Default:     "",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"name"},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_X509Subject(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "X509Subject Full X509 name specification",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"organizations": {
						SchemaProps: spec.SchemaProps{
							Description: "Organizations to be used on the Certificate.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"countries": {
						SchemaProps: spec.SchemaProps{
							Description: "Countries to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"organizationalUnits": {
						SchemaProps: spec.SchemaProps{
							Description: "Organizational Units to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"localities": {
						SchemaProps: spec.SchemaProps{
							Description: "Cities to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"provinces": {
						SchemaProps: spec.SchemaProps{
							Description: "State/Provinces to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"streetAddresses": {
						SchemaProps: spec.SchemaProps{
							Description: "Street addresses to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"postalCodes": {
						SchemaProps: spec.SchemaProps{
							Description: "Postal codes to be used on the CertificateSpec.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Default: "",
										Type:    []string{"string"},
										Format:  "",
									},
								},
							},
						},
					},
					"serialNumber": {
						SchemaProps: spec.SchemaProps{
							Description: "Serial number to be used on the CertificateSpec.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
	}
}

func schema_kmodulesxyz_client_go_api_v1_stringSetMerger(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
			},
		},
	}
}
