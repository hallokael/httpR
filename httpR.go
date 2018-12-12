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
}

func StatusOK(params ...string) HttpRun {
	return CommonRequest(http.StatusOK, params...)
}
func StatusInternalServerError(params ...string) HttpRun {
	return CommonRequest(http.StatusInternalServerError, params...)
}
func StatusBadRequest(params ...string) HttpRun {
	return CommonRequest(http.StatusBadRequest, params...)
}
func CommonRequest(code int, param ...string) HttpRun {
	line, fileName := GetCallerMessage()
	return HttpRun{
		Line:     line,
		FileName: fileName,
		Code:     code,
		Params:   param,
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
