package handlers

import (
	"fmt"
	"net/http"

	"github.com/crystaldust/configcomb/utils"
	"k8s.io/client-go/rest"
)

var istioRestClient *rest.RESTClient

func init() {
	var err error
	istioRestClient, err = utils.CreateRestClient("apis", "networking.istio.io", "v1apha3")
	if err != nil {
		panic("Failed to init rest client for istio")
	}
}

func ListVirtualServices(w http.ResponseWriter, r *http.Request) {
	req := istioRestClient.Get()
	req.SubResource("virtualservices")
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

func ListDestinationRules(w http.ResponseWriter, r *http.Request) {
	req := istioRestClient.Get()
	req.SubResource("destinationrules")
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
