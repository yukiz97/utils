package utils

//GetTypeOf return a string is indentify of data type
func GetTypeOf(input interface{}) string {
	switch input.(type) {
	case int:
		return "int"
	case string:
		return "string"
	default:
		return "unknown"
	}
}
