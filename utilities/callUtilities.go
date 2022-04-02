package utilities

import (
	"fmt"
	"io"
	"net/http"
)

var address string = ConnectConfiguration.Protocol + "://" + ConnectConfiguration.Hostname[0]

func DoCallByHost(method, hostPath string, body io.Reader) (*http.Response, error) {

	URL := ConnectConfiguration.Protocol + "://" + hostPath
	return doCall(method, URL, body)
}

func DoCallByPath(method, path string, body io.Reader) (*http.Response, error) {

	URL := address + path
	return doCall(method, URL, body)
}

func doCall(method, URL string, body io.Reader) (*http.Response, error) {

	client := createClient()

	request, err := http.NewRequest(method, URL, body)
	if err != nil {
		fmt.Printf("Creation of request failed with error %s\n", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}

	return response, nil
}
