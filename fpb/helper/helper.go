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
