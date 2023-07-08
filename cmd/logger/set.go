package logger

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

var inSetLevel string
var defaultLevel = "INFO" // TODO set this from the config file intstead than hardcoding

var LoggerSetCmd = &cobra.Command{
	Use:   "set [flags] logger_name",
	Short: "sets the log level set for a logger",
	Long: "Allows to set the log level set for a specific logger or connector plugin at runtime;" +
		"\tallowed log levels are as in Java [\"OFF\",\"FATAL\",\"ERROR\",\"WARN\",\"INFO\",\"DEBUG\",\"TRACE\"]",
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		pluginClass = args[0]
		setLevel := strings.ToUpper(inSetLevel)
		fmt.Println("input set level is", inSetLevel)
		fmt.Println("checked set level is", setLevel)
		_, err := validateLogLevel(setLevel)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		//fmt.Println("pluiginClass is " + pluginClass) // control statement print
		for _, host := range utilities.ConnectConfiguration.Hostnames {
			var loggerListURL string = host + "/admin/loggers/" + pluginClass
			fmt.Println("--- Setting Log Level", setLevel, "for Connect worker at", host, "---")
			//fmt.Println("making a call to", loggerListURL) // control statement print
			response, err := utilities.DoCallByHost(http.MethodPut, loggerListURL, bytes.NewBuffer(buildSetPayload(setLevel)))
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				utilities.PrintResponseJson(response)
			}
		}
	},
}

func init() {
	LoggerSetCmd.Flags().StringVarP(&inSetLevel, "level", "", defaultLevel, "log level to set")
}

func buildSetPayload(level string) []byte {
	payload, err := json.Marshal(map[string]interface{}{"level": level})
	if err != nil {
		fmt.Printf("JSON build failed with error %s\n", err)
	}
	return payload
}

func validateLogLevel(level string) (bool, error) {
	switch level {
	case
		"OFF",
		"FATAL",
		"ERROR",
		"WARN",
		"INFO",
		"DEBUG",
		"TRACE":
		return true, nil
	}
	return false, errors.New("invalid log level; valid log levels are [\"OFF\",\"FATAL\",\"ERROR\",\"WARN\",\"INFO\",\"DEBUG\",\"TRACE\"]")
}
