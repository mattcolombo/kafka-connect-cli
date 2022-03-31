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
