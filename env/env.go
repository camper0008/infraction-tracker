package env

import "os"

func EnvironmentVariablesSpecified() {
	if os.Getenv("DOMAIN") == "" {
		panic("DOMAIN environment variable not set")
	}

	if os.Getenv("API_KEY") == "" {
		panic("API_KEY environment variable not set")
	}
}

func Domain() string {
	return os.Getenv("DOMAIN")
}

func ApiKey() string {
	return os.Getenv("API_KEY")
}
