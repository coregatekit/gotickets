package common

// Base http response struct
type Response struct {
	Code    int    `json:"code"`    // custom code response from the server
	Message string `json:"message"` // custom message response from the server
	Data    any    `json:"data"`    // data response from the server
}
