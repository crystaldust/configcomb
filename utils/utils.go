package utils

import (
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func CreateRestClient(apiPath, group, version string) (*rest.RESTClient, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		config, err = clientcmd.BuildConfigFromFlags("", "/home/lance/.kube/config")
		if err != nil {
			return nil, err
		}
	}

	config.APIPath = apiPath
	config.GroupVersion = &schema.GroupVersion{
		Group:   group,
		Version: version,
	}
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: serializer.NewCodecFactory(runtime.NewScheme())}

	k8sRestClient, err := rest.RESTClientFor(config)
	if err != nil {
		return nil, err
	}
	return k8sRestClient, nil
}
