package model

import (
	"istio.io/api/mixer/v1/config/client"
	"istio.io/istio/mixer/adapter/memquota/config"
)

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

type StructRule struct {
	k8sResBase
	Metadata struct {
		ClusterName     string `json:"clusterName"`
		Generation      int    `json:"generation"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		Actions []*StructAction `json:"actions"`
	} `json:"spec"`
}
type StructAction struct {
	Handler   string   `json:"handler"`
	Instances []string `json:"instances"`
}

type StructQuotaSpec struct {
	k8sResBase
	Metadata struct {
		ClusterName     string `json:"clusterName"`
		Generation      int    `json:"generation"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec struct {
		Rules []*client.QuotaRule `json:"rules"`
	} `json:"spec"`
}

type StructQuotaSpecBinding struct {
	k8sResBase
	Metadata struct {
		ClusterName     string `json:"clusterName"`
		Generation      int    `json:"generation"`
		Name            string `json:"name"`
		Namespace       string `json:"namespace"`
		ResourceVersion string `json:"resourceVersion"`
	} `json:"metadata"`
	Spec *client.QuotaSpecBinding `json:"spec"`
}

func Rule(name, namespace string, actions []*StructAction) *StructRule {
	rule := &StructRule{}

	rule.APIVersion = "config.istio.io/v1alpha2"
	rule.Kind = "rule"

	rule.Metadata.Name = name
	rule.Metadata.Namespace = namespace

	rule.Spec.Actions = actions

	return rule
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

func QuotaSpec(name, namespace string, rules []*client.QuotaRule) *StructQuotaSpec {
	quotaspec := &StructQuotaSpec{}

	quotaspec.APIVersion = "config.istio.io/v1alpha2"
	quotaspec.Kind = "QuotaSpec"

	quotaspec.Metadata.Name = name
	quotaspec.Metadata.Namespace = namespace

	quotaspec.Spec.Rules = rules

	return quotaspec
}

func QuotaSpecBinding(name, namespace string, spec *client.QuotaSpecBinding) *StructQuotaSpecBinding {
	quotaspecBinding := &StructQuotaSpecBinding{}

	quotaspecBinding.APIVersion = "config.istio.io/v1alpha2"
	quotaspecBinding.Kind = "QuotaSpecBinding"

	quotaspecBinding.Metadata.Name = name
	quotaspecBinding.Metadata.Namespace = namespace

	quotaspecBinding.Spec = spec

	return quotaspecBinding
}
