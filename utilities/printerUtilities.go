package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

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
