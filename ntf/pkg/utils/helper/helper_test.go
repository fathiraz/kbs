package helper

import "testing"

func TestStingToUint64(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult uint64
	}{
		{
			name: "Success Case",
			args: args{
				str: "1",
			},
			wantResult: uint64(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := StringToUint64(tt.args.str); gotResult != tt.wantResult {
				t.Errorf("StringToUint64() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
