package sms_vendor

import (
	"ntf/config/environment"
	"testing"
)

func TestParseSmsVendorConfig(t *testing.T) {
	type fields struct {
		Env environment.Environment
	}

	tests := []struct {
		name   string
		fields fields
	}{
		{
			name: "Success Case" +
				" yaml",
			fields: fields{Env: environment.Environment{
				SmsConfigType: SmsTypeYaml,
				SmsConfigPath: ".",
			}},
		},
		{
			name: "Success Case" +
				" toml",
			fields: fields{Env: environment.Environment{
				SmsConfigType: SmsTypeToml,
				SmsConfigPath: ".",
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			environment.SetEnv(tt.fields.Env)
			ParseSmsVendorConfig()
		})
	}
}

func TestGetSmsVendorConfig(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Success Case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetSmsVendorConfig()
		})
	}
}

func TestSetSmsVendorConfig(t *testing.T) {
	type args struct {
		newSmsvendor map[string]SmsVendorConfig
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Success Case",
			args: args{newSmsvendor: map[string]SmsVendorConfig{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetSmsVendorConfig(tt.args.newSmsvendor)
		})
	}
}
