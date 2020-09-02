package utils

import "encoding/base64"

//EncodeStringToBase64 encode a string to base64 string
func EncodeStringToBase64(strInput string) string {
	strBase64 := base64.StdEncoding.EncodeToString([]byte(strInput))

	return strBase64
}
