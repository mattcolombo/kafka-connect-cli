package utilities

type ConfigurationYaml struct {
	Hostnames  []string   `yaml:"hostnames"`
	Protocol   string     `yaml:"protocol"`
	Tls        Tls        `yaml:"tls"`
	BasicAuth  BasicAuth  `yaml:"basicauth"`
	TokenAuth  TokenAuth  `yaml:"tokenauth"`
	ApiKeyAuth ApiKeyAuth `yaml:"apikeyauth"`
}

type Tls struct {
	Enabled  bool   `yaml:"enabled"`
	CaPath   string `yaml:"capath"`
	CertPath string `yaml:"certpath"`
	KeyPath  string `yaml:"keypath"`
}

type BasicAuth struct {
	Enabled bool   `yaml:"enabled"`
	User    string `yaml:"user"`
	PassRef string `yaml:"passref"`
}

type TokenAuth struct {
	Enabled  bool   `yaml:"enabled"`
	TokenRef string `yaml:"tokenref"`
}

type ApiKeyAuth struct {
	Enabled bool   `yaml:"enabled"`
	Header  string `yaml:"header"`
	Keyref  string `yaml:"keyref"`
}

type JsonError struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}

type Version struct {
	Major      string `json:"major"`
	Minor      string `json:"minor"`
	GitVersion string `json:"gitVersion"`
	GitCommit  string `json:"gitCommit"`
	BuildDate  string `json:"buildDate"`
	GoVersion  string `json:"goVersion"`
}
