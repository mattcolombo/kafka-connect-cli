package utilities

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"os"
)

func createTlsClient(capath string, certPath string, keyPath string) *http.Client {

	// the CertPool wants to add a root as a []byte so we read the file ourselves
	caCert, err := os.ReadFile(capath)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	// LoadX509KeyPair reads files, so we give it the paths
	clientCert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := tls.Config{
		RootCAs:      pool,
		Certificates: []tls.Certificate{clientCert},
	}
	transport := http.Transport{
		TLSClientConfig: &tlsConfig,
	}
	client := http.Client{
		Transport: &transport,
	}

	return &client
}

func createClient() *http.Client {
	//fmt.Println("I am creating the connect client") // control statement print
	if ConnectConfiguration.Tls.Enabled {
		return createTlsClient(ConnectConfiguration.Tls.CaPath, ConnectConfiguration.Tls.CertPath, ConnectConfiguration.Tls.KeyPath)
	} else {
		return &http.Client{}
	}
}
