package model

import "istio.io/istio/mixer/adapter/memquota/config"

type StructMemQuota struct {
	k8sResBase
	Metadata struct {
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		Quotas []*config.Params_Quota `json:"quotas"`
	} `json:"spec"`
}

// type Jsonable interface {
//     Json() string
// }

type k8sResBase struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
}

type StructQuota struct {
	k8sResBase
	Metadata struct {
		ClusterName     string `json:"clusterName"`
		Generation      int    `json:"generation"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		Dimensions map[string]string `json:"dimensions"`
	} `json:"spec"`
}

func MemQuota(name, namespace string, quotas []*config.Params_Quota) *StructMemQuota {
	memquota := &StructMemQuota{}

	memquota.APIVersion = "config.istio.io/v1alpha2"
	memquota.Kind = "memquota"

	memquota.Metadata.Name = name
	memquota.Metadata.Namespace = namespace

	memquota.Spec.Quotas = quotas

	return memquota
}

func Quota(name, namespace string, dimensions map[string]string) *StructQuota {
	quota := &StructQuota{}

	quota.APIVersion = "config.istio.io/v1alpha2"
	quota.Kind = "quota"

	quota.Metadata.Name = name
	quota.Metadata.Namespace = namespace

	quota.Spec.Dimensions = dimensions

	return quota
}
