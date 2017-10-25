/*
Copyright 2017 the project contributors.

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

package main

import (
	"fmt"
	"os"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
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
