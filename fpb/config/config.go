package config

import (
	"fpb/helper"
	"github.com/joho/godotenv"
	"os"
)

type (

	// Environment struct to hold env data
	Environment struct {
		DefaultOwner  string
		DefaultApples uint64
		DefaultCakes  uint64
	}
)

var env Environment

// ParseEnv function to get env data
func ParseEnv() {

	// read .env
	godotenv.Load()

	env = Environment{
		DefaultOwner:  os.Getenv("DEFAULT_OWNER"),
		DefaultApples: helper.StingToUint64(os.Getenv("DEFAULT_APPLES")),
		DefaultCakes:  helper.StingToUint64(os.Getenv("DEFAULT_CAKES")),
	}
}

// GetEnv function to get env data
func GetEnv() Environment {
	return env
}
