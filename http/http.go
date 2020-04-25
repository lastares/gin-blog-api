package http

const ErrorCode = -1;

var statusText = map[string]string{
	"Params error.": "参数错误",
	"Tag does not exist.": "标签不存在",
	"The tag already exists.": "标签已经存在",
	"Deleted failed.": "删除失败",
	"Created failed.": "添加失败",
	"Updated failed.": "更新失败",
}

func Translate(errorMessage string) string {
	return statusText[errorMessage]
}

