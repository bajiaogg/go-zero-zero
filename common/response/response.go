package response

import "net/http"

type RespSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}

type RespError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"message"`
}

func Success(data interface{}) *RespSuccess {
	return &RespSuccess{http.StatusOK, "SUCCESS", data}
}

func Error(errCode uint32, errMsg string) *RespError {
	return &RespError{errCode, errMsg}
}
