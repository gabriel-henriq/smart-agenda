package models

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func ResponseData(code, message string, success bool, data interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Success: success,
		Data:    data,
	}
}
