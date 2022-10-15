package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func PrintEmptyBodyResponse(response *http.Response, successCode int, message string) {
	defer response.Body.Close()

	if response.StatusCode == successCode {
		fmt.Println(message)
		//fmt.Println("Connect responds:", response.Status) // control statement not quite required in the future
	} else {
		fmt.Println("Connect responds:", response.Status, "-", extractMessageFromJsonError(response))
	}
}

func PrintResponseJson(response *http.Response) {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	PrettyPrintJson(body)
}

func PrettyPrintJson(data []byte) {
	var prettyData bytes.Buffer
	json.Indent(&prettyData, data, "", "  ")
	fmt.Println(prettyData.String())
}

func extractMessageFromJsonError(response *http.Response) string {
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := JsonError{}
	json.Unmarshal(body, &data)
	return data.Message
}
