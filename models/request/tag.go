package request

type TagCreateModel struct {
	TagName   string `json:"tagName" validate:"required" label:"标签名称"`
	TagStatus int    `json:"tagStatus" validate:"required,oneof=10 20" label:"标签状态"`
}
