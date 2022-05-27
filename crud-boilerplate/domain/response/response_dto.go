package response

func SuccessResponseMap(data interface{}, message string) *Response {
	return &Response{
		Status:  "success",
		Message: message,
		Data:    data,
	}
}
