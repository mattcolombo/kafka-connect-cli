package utilities

type Configuration struct {
	Protocol  string
	Hostname  []string
	Operation string
	TlsEnable bool
	CaPath    string
	CertPath  string
	KeyPath   string
}

type JsonError struct {
	Code    int    `json:"error_code"`
	Message string `json:"message"`
}
