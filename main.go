package main

import (
	"context"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	// Carregar configurações do Kubernetes a partir do arquivo kubeconfig
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	rules.ExplicitPath = "kubeconfig.yaml"
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		rules,
		&clientcmd.ConfigOverrides{},
	)
	
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		panic(err.Error())
	}
	
	// Criar o cliente Kubernetes
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	
	// Listar os pods do namespace padrão
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
	}
	
	// Imprimir os nomes dos pods
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
}
