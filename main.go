package main

import (
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/pkg/api/v1"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	loader := clientcmd.NewDefaultClientConfigLoadingRules()
	clientConfig, err := clientcmd.BuildConfigFromKubeconfigGetter("", loader.Load)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating client config: %v\n", err)
		os.Exit(1)
	}

	client := kubernetes.NewForConfigOrDie(clientConfig)

	ns := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: "foo-",
		},
	}

	updated, err := client.CoreV1().Namespaces().Create(ns)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating namespace: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("namespace created: %#v\n", updated)
}
