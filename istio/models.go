package istio

import "istio.io/api/mixer/v1/config/client"

// The instance for the 'rules' field in QuotaSpec
func createQuotaRule() []*client.QuotaRule {
	quota := &client.Quota{
		Quota:  "",
		Charge: 1,
	}
	quotaRule := &client.QuotaRule{
		Quotas: []*client.Quota{quota},
	}

	return []*client.QuotaRule{quotaRule}
}

func createQuotaSpec(rules []*client.QuotaRule) *client.QuotaSpec {
	quotaSpec := &client.QuotaSpec{
		Rules: rules,
	}

	return quotaSpec
}

func createQuotaSpecBinding() *client.QuotaSpecBinding {
	services := []*client.IstioService{
		{
			Name:      SERVICE_NAME,
			Namespace: "default",
		},
	}
	quotaSpecRefs := []*client.QuotaSpecBinding_QuotaSpecReference{
		{
			Name:      "",
			Namespace: "default",
		},
	}

	return &client.QuotaSpecBinding{
		Services:   services,
		QuotaSpecs: quotaSpecRefs,
	}
}
