package main

import (
	"context"
	"log"
	"os"

	v1 "get.porter.sh/operator/api/v1"
	"get.porter.sh/operator/clientset/v1/porterclientset"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	ctx := context.Background()
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	porterClient, err := porterclientset.NewPorterClientSet(config)
	if err != nil {
		log.Fatal("error getting the porter client")
	}
	porterList := &v1.InstallationList{}
	err = porterClient.List(ctx, porterList)
	if err != nil {
		log.Fatal("error getting the list of installations")
	}
	for _, install := range porterList.Items {
		log.Printf("installation: %s -- %s", install.Name, install.Namespace)
	}
	if len(porterList.Items) < 1 {
		log.Print("there are no porter installations on this cluster")
	}
	log.Print("done getting the installations")
}
