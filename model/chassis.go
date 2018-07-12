package model

type Chassis struct {
	APPLICATIONID     string  `yaml:"APPLICATION_ID"`
	ManagementAbility string  `yaml:"managementAbility"`
	Cse               Cse     `yaml:"cse"`
	Tracing           Tracing `yaml:"tracing"`
}

type Cse struct {
	Service        Service        `yaml:"service"`
	Protocols      Protocols      `yaml:"protocols"`
	Handler        Handler        `yaml:"handler"`
	Flowcontrol    Flowcontrol    `yaml:"flowcontrol"`
	CircuitBreaker CircuitBreaker `yaml:"circuitBreaker"`
	Fallbackpolicy Fallbackpolicy `yaml:"fallbackpolicy"`
	Metrics        Metrics        `yaml:"metrics"`
}

type Tracing struct {
	Enabled         bool   `yaml:"enabled"`
	CollectorType   string `yaml:"collectorType"`
	CollectorTarget string `yaml:"collectorTarget"`
}

type Service struct {
	Registry struct {
		Type    string `yaml:"type"`
		Address string `yaml:"address"`
	} `yaml:"registry"`
}

type Protocols struct {
	Rest struct {
		ListenAddress    string `yaml:"listenAddress"`
		AdvertiseAddress string `yaml:"advertiseAddress"`
	} `yaml:"rest"`
}

type Handler struct {
	Chain struct {
		Provider struct {
			Default string `yaml:"default"`
		} `yaml:"Provider"`
	} `yaml:"chain"`
}
type Flowcontrol struct {
	Provider struct {
		QPS struct {
			Enabled bool `yaml:"enabled"`
			Global  struct {
				Limit int `yaml:"limit"`
			} `yaml:"global"`
			Limit map[string]int `yaml:"limit"`
		} `yaml:"qps"`
	} `yaml:"Provider"`
}
type CircuitBreaker struct {
	Provider struct {
		Enabled                   bool `yaml:"enabled"`
		RequestVolumeThreshold    int  `yaml:"requestVolumeThreshold"`
		SleepWindowInMilliseconds int  `yaml:"sleepWindowInMilliseconds"`
	} `yaml:"Provider"`
}
type Fallbackpolicy struct {
	Provider struct {
		Policy string `yaml:"policy"`
	} `yaml:"Provider"`
}
type Metrics struct {
	APIPath                string `yaml:"apiPath"`
	Enable                 bool   `yaml:"enable"`
	EnableGoRuntimeMetrics bool   `yaml:"enableGoRuntimeMetrics"`
}
