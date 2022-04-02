package utilities

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// TODO move this somewhere more logical (if even required)
func PrettyPrint(data []byte) {
	var prettyData bytes.Buffer
	json.Indent(&prettyData, data, "", "  ")
	fmt.Println(prettyData.String())
}
