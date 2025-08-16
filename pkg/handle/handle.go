package handle

import (
	"go-web-template/logger"
	"go-web-template/pkg/buserr"
	"go-web-template/pkg/dto"
)

// Response 通用响应结构，用于 Swagger 文档
type Response = dto.Result

func Result(code int, result interface{}, msg string) dto.Result {
	return dto.Result{
		ErrCode: code,
		ErrMsg:  msg,
		Result:  result,
	}
}

func Ok() dto.Result {
	return Result(buserr.Ok, nil, "ok")
}
func OkWithMsg(msg string) dto.Result {
	return Result(buserr.Ok, nil, msg)
}

func Success(result interface{}) dto.Result {
	return Result(buserr.Ok, result, "ok")
}

func Error(error string) dto.Result {
	return dto.Result{
		ErrCode: buserr.Err,
		ErrMsg:  error,
	}
}

func Fail(code int, msg string) dto.Result {
	return Result(code, nil, msg)
}

func FailWithMsg(code int, msg string) dto.Result {
	return Result(code, nil, msg)
}

func HandleError(err error, codes ...interface{}) {
	if err != nil {
		if len(codes) > 0 {
			logger.Error("%v\n%s", codes, err.Error())
		} else {
			// panic(err.Error())
		}
	}
}
