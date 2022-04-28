package helper

import (
	"strconv"
)

// StingToUint64 to set string to uint64
func StingToUint64(str string) (result uint64) {
	convert, err := strconv.Atoi(str)
	if err == nil {
		result = uint64(convert)
	}

	return
}
