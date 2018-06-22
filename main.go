package main

import (
	"fmt"
	"net/http"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var k8sRestClient *rest.RESTClient
var err error

func main() {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "/home/lance/.kube/config")
		if err != nil {
			panic(err.Error())
		}
	}

	config.APIPath = "/apis"
	config.GroupVersion = &schema.GroupVersion{
		Group:   "networking.istio.io",
		Version: "v1alpha3",
	}
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}

	go getPodNum(config)

	k8sRestClient, err = rest.RESTClientFor(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/virtualservices", getVirtualServices)
	mux.HandleFunc("/destinationrules", getDestinationRules)
	mux.HandleFunc("/deployments", getDeployments)
	if err := http.ListenAndServe(":8882", mux); err != nil {
		fmt.Println(err)
	}

}

func getPodNum(config *rest.Config) {
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
}

func getVirtualServices(w http.ResponseWriter, r *http.Request) {

	req := k8sRestClient.Get()
	req.SubResource("virtualservices")
	result := req.Do()
	rawBody, err := result.Raw()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	} else {
		w.Write(rawBody)
		return
	}

}

func getDeployments(w http.ResponseWriter, r *http.Request) {
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
		fmt.Println(err)
		return
	}

	req := k8sRestClient.Get()
	req.SubResource("deployments")
	result := req.Do()
	rawBody, err := result.Raw()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	} else {
		w.Write(rawBody)
		return
	}
}
func getDestinationRules(w http.ResponseWriter, r *http.Request) {
	req := k8sRestClient.Get()
	req.SubResource("destinationrules")
	result := req.Do()
	rawBody, err := result.Raw()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, err.Error())
		return
	} else {
		w.Write(rawBody)
		return
	}
}
