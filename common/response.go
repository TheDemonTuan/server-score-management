package common

type Response struct {
	StatusCode int         `json:"code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data"`
}
