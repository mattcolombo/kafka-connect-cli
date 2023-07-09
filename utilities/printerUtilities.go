package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

func PrintEmptyBodyResponse(response *http.Response, successCode int, message string) {
	defer response.Body.Close()

	if response.StatusCode == successCode {
		fmt.Println(message)
		//fmt.Println("Connect responds:", response.Status) // control statement not quite required in the future
	} else {
		jsonError, err := extractMessageFromJsonError(response)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unable to extractMessageFromJsonError %s", jsonError)
			os.Exit(1)
		}
		fmt.Println("Connect responds:", response.Status, "-")
	}
}

func PrintResponseJson(response *http.Response) {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err = PrettyPrintJson(body); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func PrettyPrintJson(data []byte) error {
	var prettyData bytes.Buffer
	if err := json.Indent(&prettyData, data, "", "  "); err != nil {
		return err
	}
	fmt.Println(prettyData.String())
	return nil
}

func PrettyPrintConfigYaml(yamlData ConfigurationYaml) {
	byte, err := yaml.Marshal(&yamlData)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", string(byte))
}

func extractMessageFromJsonError(response *http.Response) (string, error) {
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := JsonError{}
	if err = json.Unmarshal(body, &data); err != nil {
		return "", err
	}
	return data.Message, nil
}
