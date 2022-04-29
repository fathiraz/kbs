package result

import "net/http"

type (

	// Result struct to set data json result
	Result struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Error   interface{} `json:"error,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}
)

// ResultSuccess function to show json result success
func ResultSuccess(message string, data interface{}) (result *Result) {
	result = &Result{
		Success: true,
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	}
	return
}

// ResultError function to show json result error
func ResultError(message string, err error) (result *Result) {
	result = &Result{
		Success: false,
		Code:    http.StatusBadRequest,
		Message: message,
		Error:   err,
	}
	return
}
