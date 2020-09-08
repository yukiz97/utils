package utils

import (
	"fmt"
	"strings"
)

//IntArrayToString join array to string with
func IntArrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}
