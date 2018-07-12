package handlers

import (
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/juzhen/k8s-client-test/utils"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var k8sRestClient *rest.RESTClient

func init() {
	var err error
	k8sRestClient, err = utils.CreateRestClient("apis", "apps", "v1")
	if err != nil {
		panic("Failed to init rest client for k8s resources")
	}
}

func GetPodNum(config *rest.Config) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		return
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func GetDeployments(w http.ResponseWriter, r *http.Request) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "/home/lance/.kube/config")
		if err != nil {
			panic(err.Error())
		}
	}

	config.APIPath = "/apis"
	config.GroupVersion = &schema.GroupVersion{
		Group:   "apps",
		Version: "v1",
	}
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}

	k8sRestClient, err := rest.RESTClientFor(config)
	if err != nil {
		return
	}

	req := k8sRestClient.Get()
	req.SubResource("deployments")
	result := req.Do()
	rawBody, err := result.Raw()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	} else {
		w.Write(rawBody)
		return
	}
}
