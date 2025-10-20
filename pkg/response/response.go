package resp

import (
	"github.com/rPniu/all/pkg/e"
	"time"
)

type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp int64       `json:"timestamp"`
}

func Success(data interface{}) Response {
	resp := Response{
		Code:      1000,
		Message:   "success",
		Data:      data,
		Timestamp: time.Now().Unix(),
	}
	return resp
}

func Failure(code int) Response {
	resp := Response{
		Code:      code,
		Message:   e.GetMsg(code),
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
	return resp
}

func FailureAddDetail(code int, detail string) Response {
	resp := Response{
		Code:      code,
		Message:   e.GetMsg(code) + "detail:" + detail,
		Data:      nil,
		Timestamp: time.Now().Unix(),
	}
	return resp
}
