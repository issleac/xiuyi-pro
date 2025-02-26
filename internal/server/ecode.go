package server

import "github.com/gin-gonic/gin"

// BuildDataResponse 构建包含错误码和错误信息的响应
func BuildDataResponse(ctx *gin.Context, data interface{}) map[string]interface{} {
	response := make(map[string]interface{})
	requestId := ctx.GetString("request_id")
	response["request_id"] = requestId
	response["code"] = 0
	response["message"] = "ok"
	response["data"] = data
	return response
}

// BuildErrorResponse 构建包含错误码和错误信息的响应
func BuildErrorResponse(ctx *gin.Context, errCode int64) map[string]interface{} {
	response := make(map[string]interface{})
	requestId := ctx.GetString("request_id")
	response["request_id"] = requestId
	response["code"] = errCode
	message := GetErrMsg(errCode)
	response["message"] = message
	return response
}

func GetErrMsg(code int64) string {
	return "出错了" //todo:errCode
	//return codeMsg[code]
}
