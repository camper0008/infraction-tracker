package env

import (
	"fmt"
	"os"
)

func EnvironmentVariablesSpecified() {
	required := []string{"DOMAIN", "API_KEY"}

	for i := range required {
		if os.Getenv(required[i]) == "" {
			panic(fmt.Sprintf("%s environment variable not set", required[i]))
		}

	}
}

func Domain() string {
	return os.Getenv("DOMAIN")
}

func ApiKey() string {
	return os.Getenv("API_KEY")
}
