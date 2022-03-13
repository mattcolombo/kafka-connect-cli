package utilities

type Configuration struct {
	Hostname  string
	Operation string
	TlsEnable bool
	CaPath    string
	CertPath  string
	KeyPath   string
}
