package httpR

import (
	"net/http"
	"runtime"
	"strings"
)

type HttpRun struct {
	Line     int
	FileName string
	Code     int
	Params   []string
	Resp     interface{}
}

func StatusOK(resp interface{}, params ...string) HttpRun {
	return CommonRequest(http.StatusOK, resp, params...)
}
func StatusInternalServerError(resp interface{}, params ...string) HttpRun {
	return CommonRequest(http.StatusInternalServerError, resp, params...)
}
func StatusBadRequest(resp interface{}, params ...string) HttpRun {
	return CommonRequest(http.StatusBadRequest, resp, params...)
}
func CommonRequest(code int, resp interface{}, param ...string) HttpRun {
	line, fileName := GetCallerMessage()
	return HttpRun{
		Line:     line,
		FileName: fileName,
		Code:     code,
		Params:   param,
		Resp:     resp,
	}
}
func GetCallerMessage() (int, string) {
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "CantGetfileName"
	}
	fileNameAll := strings.Split(file, "/")
	fileName := fileNameAll[len(fileNameAll)-1]
	return line, fileName
}
