package result

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestResultSuccess(t *testing.T) {
	type args struct {
		message string
		data    interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult *Result
	}{
		{
			name: "Success Case",
			args: args{
				message: "test",
				data:    "test",
			},
			wantResult: &Result{
				Success: true,
				Code:    http.StatusOK,
				Message: "test",
				Data:    "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ResultSuccess(tt.args.message, tt.args.data); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ResultSuccess() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestResultError(t *testing.T) {
	type args struct {
		message string
		err     error
	}
	tests := []struct {
		name       string
		args       args
		wantResult *Result
	}{
		{
			name: "Success Case",
			args: args{
				message: "test",
				err:     fmt.Errorf("test"),
			},
			wantResult: &Result{
				Success: false,
				Code:    http.StatusBadRequest,
				Message: "test",
				Error:   fmt.Errorf("test"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := ResultError(tt.args.message, tt.args.err); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("ResultError() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
