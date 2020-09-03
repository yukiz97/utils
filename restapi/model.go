package restapi

//ResultMessage return a struct with {"result" : "return message"}
type IDItemAndMessage struct {
	ID      interface{} `json:"id"`
	Message string      `json:"message"`
}
