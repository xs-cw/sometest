package response

type BaseResponse struct {
	Code    int
	Data    interface{}
	Message string
}

// Success ...
func Success(data interface{}) (int, BaseResponse) {
	return 200, BaseResponse{
		Code:    200,
		Data:    data,
		Message: "Success",
	}
}

// Failed ...
func Failed(code int, msg string) (int, interface{}) {
	return 500, BaseResponse{
		Code:    code,
		Data:    nil,
		Message: msg,
	}
}
