package main

import (
	"fmt"
	"net/http"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
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

	k8sRestClient, err = rest.RESTClientFor(config)
	if err != nil {
		fmt.Println(err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/virtualservices", getVirtualServices)
	mux.HandleFunc("/destinationrules", getDestinationRules)
	http.ListenAndServe(":8882", mux)
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
