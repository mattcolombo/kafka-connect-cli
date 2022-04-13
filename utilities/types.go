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
	Certpath string `yaml:"certpath"`
	Keypath  string `yaml:"keypath"`
}

type BasicAuth struct {
	Enabled bool   `yaml:"enabled"`
	User    string `yaml:"user"`
	Passref string `yaml:"passref"`
}

type TokenAuth struct {
	Enabled  bool   `yaml:"enabled"`
	Tokenref string `yaml:"tokenref"`
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
