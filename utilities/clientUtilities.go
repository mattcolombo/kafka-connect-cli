package utilities

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func createTlsClient(capath string, certpath string, keypath string) *http.Client {

	// the CertPool wants to add a root as a []byte so we read the file ourselves
	caCert, err := ioutil.ReadFile(capath)
	if err != nil {
		log.Fatal(err)
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caCert)

	// LoadX509KeyPair reads files, so we give it the paths
	clientCert, err := tls.LoadX509KeyPair(certpath, keypath)
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
	fmt.Println("I am creating the connect client") // control statement print - TOREMOVE
	if ConnectConfiguration.TlsEnable {
		return createTlsClient(ConnectConfiguration.CaPath, ConnectConfiguration.CertPath, ConnectConfiguration.KeyPath)
	} else {
		return &http.Client{}
	}
}
