package environment

import (
	"github.com/joho/godotenv"
	"ntf/pkg/utils/helper"
	"os"
)

type (

	// Environment struct to hold env data
	Environment struct {
		AppName       string
		AppPort       string
		AppDebug      bool
		BasicUsername string
		BasicPassword string
		SmsConfigType string
		SmsConfigPath string
	}
)

var env Environment

// ParseEnv function to get env data
func ParseEnv() {

	// read .env
	godotenv.Load()

	env = Environment{
		AppName:       os.Getenv("APP_NAME"),
		AppPort:       os.Getenv("APP_PORT"),
		AppDebug:      helper.StringToBool(os.Getenv("APP_DEBUG")),
		BasicUsername: os.Getenv("BASIC_USERNAME"),
		BasicPassword: os.Getenv("BASIC_PASSWORD"),
		SmsConfigType: os.Getenv("SMS_CONFIG_TYPE"),
		SmsConfigPath: os.Getenv("SMS_CONFIG_PATH"),
	}
}

// GetEnv function to get env data
func GetEnv() Environment {
	return env
}

// SetEnv function to set env data
func SetEnv(newEnv Environment) {
	env = newEnv
}
