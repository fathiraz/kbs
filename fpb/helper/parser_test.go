package helper

import (
	"reflect"
	"testing"
)

func TestParseFlag(t *testing.T) {
	type args struct {
		defaultParser *Parser
	}
	tests := []struct {
		name       string
		args       args
		wantParser Parser
	}{
		{
			name: "Success Case",
			args: args{
				defaultParser: &Parser{
					Owner:  "test",
					Cakes:  1,
					Apples: 1,
				},
			},
			wantParser: Parser{
				Owner:  "test",
				Cakes:  1,
				Apples: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotParser := ParseFlag(tt.args.defaultParser); !reflect.DeepEqual(gotParser, tt.wantParser) {
				t.Errorf("ParseFlag() = %v, want %v", gotParser, tt.wantParser)
			}
		})
	}
}
