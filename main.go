package main

import (
	"context"
	"log"
	"os"

	v1 "get.porter.sh/operator/api/v1"
	"get.porter.sh/operator/clientset/v1/porterclientset"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	ctx := context.Background()
	kubeconfig := os.Getenv("HOME") + "/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal("failed to create rest config")
	}
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

	paramList := &v1.ParameterSetList{}
	err = porterClient.List(ctx, paramList)
	if err != nil {
		log.Fatal("error getting the list of parameter sets")
	}
	for _, paramset := range paramList.Items {
		log.Printf("parameterset: %s -- %s", paramset.Name, paramset.Namespace)
	}
	if len(paramList.Items) < 1 {
		log.Print("there are no parameter sets on this cluster")
	}
	log.Print("done getting the paramtersets")

	agentconfigList := &v1.AgentConfigList{}
	err = porterClient.List(ctx, agentconfigList)
	if err != nil {
		log.Fatal("error getting the list of agentconfigs")
	}
	for _, agentConfig := range agentconfigList.Items {
		log.Printf("agentConfig: %s -- %s", agentConfig.Name, agentConfig.Namespace)
	}
	if len(agentconfigList.Items) < 1 {
		log.Print("there are no agentconfigs on this cluster")
	}
	log.Print("done getting the agentconfigs")

	porterConfigList := &v1.PorterConfigList{}
	err = porterClient.List(ctx, porterConfigList)
	if err != nil {
		log.Fatal("error getting the list of porter configs")
	}
	for _, porterConfig := range agentconfigList.Items {
		log.Printf("porterConfig: %s -- %s", porterConfig.Name, porterConfig.Namespace)
	}
	if len(agentconfigList.Items) < 1 {
		log.Print("there are no porter configs on this cluster")
	}
	log.Print("done getting the porter configs")

	credentialsetList := &v1.CredentialSetList{}
	err = porterClient.List(ctx, credentialsetList)
	if err != nil {
		log.Fatal("error getting the list of porter configs")
	}
	for _, credset := range credentialsetList.Items {
		log.Printf("credential set: %s -- %s", credset.Name, credset.Namespace)
	}
	if len(credentialsetList.Items) < 1 {
		log.Print("there are no credential sets  on this cluster")
	}
	log.Print("done getting the credential sets")

	log.Print("creating a AgentConfig")
	porterAgentAction := &v1.AgentAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "porter-agent-action",
			Namespace: "default",
		},
		Spec: v1.AgentActionSpec{
			Command: []string{"curl", "-v", "https://porter.sh"},
		},
	}

	err = porterClient.Create(ctx, porterAgentAction)
	if err != nil {
		log.Fatal("error creating porter agent action")
	}
	gotPorterAgentAction := &v1.AgentAction{}
	err = porterClient.Get(ctx, types.NamespacedName{Name: "porter-agent-action", Namespace: "default"}, gotPorterAgentAction)
	if err != nil {
		log.Fatal("error getting porter agent action")
	}
	log.Printf("porter agent action: %+v", gotPorterAgentAction)
}
