package environment

import (
	"reflect"
	"testing"
)

func TestParseEnv(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Succxess Case",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseEnv()
		})
	}
}

func TestGetEnv(t *testing.T) {
	tests := []struct {
		name string
		want Environment
	}{
		{
			name: "Success Case",
			want: Environment{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetEnv(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}
