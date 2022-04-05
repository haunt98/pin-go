package pin

import (
	"github.com/spf13/cast"
)

// If length == 4
// 1 -> "1" -> "0001"
func pinToString(number int64, length int) string {
	if number < 0 {
		return ""
	}

	numberStr := cast.ToString(number)
	lenNumberStr := len(numberStr)
	if lenNumberStr < length {
		// Fill up "0" in prefix
		for i := 0; i < length-lenNumberStr; i++ {
			numberStr = "0" + numberStr
		}
	}

	return numberStr
}
