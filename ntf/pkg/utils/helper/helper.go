package helper

import (
	"strconv"
)

// StringToUint64 to set string to uint64
func StringToUint64(str string) (result uint64) {
	convert, err := strconv.Atoi(str)
	if err == nil {
		result = uint64(convert)
	}

	return
}

// StringToBool to set string to bool
func StringToBool(str string) (result bool) {
	convert, err := strconv.ParseBool(str)
	if err == nil {
		result = convert
	}

	return
}
