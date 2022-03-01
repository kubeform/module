package main

import (
	"kubeform.dev/module/pkg/cmds"
	"log"

	_ "go.bytebuilders.dev/license-verifier/info"
	"gomodules.xyz/logs"
	_ "k8s.io/client-go/kubernetes/fake"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/klog/v2"
)

func main() {
	rootCmd := cmds.NewRootCmd(Version)
	logs.Init(rootCmd, true)
	log.SetOutput(logs.HTTPLogger{})
	defer logs.FlushLogs()

	if err := rootCmd.Execute(); err != nil {
		klog.Fatal(err)
	}
}
