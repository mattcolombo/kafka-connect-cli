package utilities

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// defining the client here means that the client will be created when the module is loaded and will be used throughtout without being recreated for every call
var client *http.Client = createClient()
var address string = ConnectConfiguration.Protocol + "://" + ConnectConfiguration.Hostnames[0]

func DoCallByHost(method, hostPath string, body io.Reader) (*http.Response, error) {

	URL := ConnectConfiguration.Protocol + "://" + hostPath
	return doCall(method, URL, body)
}

func DoCallByPath(method, path string, body io.Reader) (*http.Response, error) {

	URL := address + path
	return doCall(method, URL, body)
}

func doCall(method, URL string, body io.Reader) (*http.Response, error) {

	request, err := http.NewRequest(method, URL, body)
	if err != nil {
		fmt.Printf("Creation of request failed with error %s\n", err)
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")

	// adding special headers for various authentication methods
	if ConnectConfiguration.BasicAuth.Enabled {
		user := ConnectConfiguration.BasicAuth.User
		pass := os.Getenv(ConnectConfiguration.BasicAuth.Passref)
		request.SetBasicAuth(user, pass)
	}
	if ConnectConfiguration.TokenAuth.Enabled {
		btoken := "Bearer " + os.Getenv(ConnectConfiguration.TokenAuth.Tokenref)
		request.Header.Add("Authorization", btoken)
	}
	if ConnectConfiguration.ApiKeyAuth.Enabled {
		header := os.Getenv(ConnectConfiguration.ApiKeyAuth.Header)
		tokenValue := os.Getenv(ConnectConfiguration.ApiKeyAuth.Keyref)
		request.Header.Add(header, tokenValue)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return nil, err
	}

	return response, nil
}
