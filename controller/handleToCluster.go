/**
* @description :
* @author : Jarick
* @Date : 2022-06-26
* @Url : http://CloudWebOps
 */
package controller

/*
Copyright 2016 The Kubernetes Authors.

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

// Note: the example only works with the code within the same release/branch.

import (
	"CloudDevKubernetes/util"
	"flag"
	"log"
	"net"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/openstack"
)

func GetConfig() (config *rest.Config, err error) {

	util.ErrPanicDebug(true)

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// creates the in-cluster config
	config, err = rest.InClusterConfig()
	if err != nil {
		// use the current context in kubeconfig
		config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
		if err != nil {
			return nil, err
			// panic(err.Error())
		}
	}
	return config, err
}

func HandleToCluster() (clientset *kubernetes.Clientset) {
	util.ErrPanicDebug(true)
	config, _ := GetConfig()
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Println(err)
		panic(err.Error())
	}
	return
}

func HandleToK8sAndMetric(debug bool) (clientset *kubernetes.Clientset, mClientset *metrics.Clientset) {

	util.ErrPanicDebug(true)
	config, _ := GetConfig()

	// 针对网络端口做探测
	conn, err := net.DialTimeout("tcp", strings.Split(config.Host, "https://")[1], util.TimeFuncDuration(1, "Second"))
	if err != nil {
		log.Println(err)
		if debug {
			panic(err.Error())
		}
		os.Exit(0)
	}
	defer conn.Close()

	clientset, err0 := kubernetes.NewForConfig(config)
	if err0 != nil {
		log.Println(err0)
		if debug {
			panic(err0.Error())
		}
	}
	mClientset, err1 := metrics.NewForConfig(config)
	if err1 != nil {
		log.Println(err1)
		if debug {
			panic(err1.Error())
		}
	}
	return
}
