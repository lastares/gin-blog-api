package http

const ErrorCode = -1

var statusText = map[string]string{
	"ok.":                     "ok",
	"Params error.":           "参数错误",
	"Tag does not exist.":     "标签不存在",
	"The tag already exists.": "标签已经存在",
	"Deleted failed.":         "删除失败",
	"Created failed.":         "添加失败",
	"Updated failed.":         "更新失败",
	"Access token failed.":    "获取token失败",
	"Token expired.":          "token 过期",
	"Token check failed.":     "Token鉴权失败",
	"Error auth token.":       "Token生成失败",
	"Token error.":            "Token错误",
}

func Translate(errorMessage string) string {
	return statusText[errorMessage]
}

type Success struct {
	Code int         `json:"code" example:"200"`
	Data interface{} `json:"data"`
}

type Ok struct {
	Code int `json:"code" example:"200"`
}

type Failed struct {
	Code int    `json:"code" example:"-1"`
	Msg  string `json:"msg" example:"failed"`
}
