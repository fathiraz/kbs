package sms_vendor

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"ntf/config/environment"
)

const (
	SmsTypeYaml = "yaml"
	SmsTypeToml = "toml"

	SmsVendorConfigFile = "sms_vendor"

	SmsVendorNameKita     = "sms_vendor_kita"
	SmsVendorNameBisa     = "sms_vendor_bisa"
	SmsVendorNameKitabisa = "sms_vendor_kitabisa"
)

type (

	// SmsVendorConfig struct to hold sms vendor data
	SmsVendorConfig struct {
		Name      string `yaml:"name" toml:"name" json:"name"`
		Url       string `yaml:"url" toml:"url" json:"url"`
		Endpoint  string `yaml:"endpoint" toml:"endpoint" json:"endpoint"`
		Token     string `yaml:"token" toml:"token" json:"token"`
		IsDefault bool   `yaml:"is_default" toml:"is_default" json:"is_default"`
		IsActive  bool   `yaml:"is_active" toml:"is_active" json:"is_active"`
		Timeout   string `yaml:"timeout" toml:"timeout" json:"timeout"`
	}
)

var smsVendorConfigs map[string]SmsVendorConfig

// ParseSmsVendorConfig function to get sms vendor data
func ParseSmsVendorConfig() {
	var (
		filename = fmt.Sprintf("%s/%s", environment.GetEnv().SmsConfigPath, SmsVendorConfigFile)
	)

	// check type
	switch environment.GetEnv().SmsConfigType {
	case SmsTypeToml:
		filename += ".toml"

		// read file
		tomlfile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// unmarshal toml
		err = toml.Unmarshal(tomlfile, &smsVendorConfigs)
		if err != nil {
			log.Fatal(err)
		}

	case SmsTypeYaml:
		filename += ".yaml"

		// read file
		yamlfile, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// unmarshal yaml
		err = yaml.Unmarshal(yamlfile, &smsVendorConfigs)
		if err != nil {
			log.Fatal(err)
		}
	}

	// validate sms vendor
	ValidateSmsVendorConfig()
}

// GetSmsVendorConfig function to get sms vendor data
func GetSmsVendorConfig() map[string]SmsVendorConfig {
	return smsVendorConfigs
}

// SetSmsVendorConfig function to set sms vendor data
func SetSmsVendorConfig(newSmsvendor map[string]SmsVendorConfig) {
	smsVendorConfigs = newSmsvendor
}

// ValidateSmsVendorConfig function to validate sms vendor data
func ValidateSmsVendorConfig() {

	var (
		countDefault, countIsActive int
	)

	for _, smsVendor := range smsVendorConfigs {
		if smsVendor.IsDefault {
			countDefault++
		}

		if smsVendor.IsActive {
			countIsActive++
		}
	}

	// default cant be 0
	if countDefault == 0 {
		log.Fatal(fmt.Errorf("sms vendor must have at least 1 vendor with default true"))
	}

	// is_active cant be 0
	if countDefault == 0 {
		log.Fatal(fmt.Errorf("sms vendor must have at least 1 vendor with is_active true"))
	}

	// must have 1 default and 1 is_active
	if countDefault+countIsActive == 0 {
		log.Fatal(fmt.Errorf("sms vendor must have at least 1 vendor with default true and is_active true"))
	}

	// default cant be more than 1
	if countDefault > 1 {
		log.Fatal(fmt.Errorf("sms vendor default cannot be more than 1"))
	}
}
