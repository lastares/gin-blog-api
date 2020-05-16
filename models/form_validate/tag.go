package form_validate

// 标签校验结构体
type TagCreateModel struct {
	TagName   string `json:"tagName" validate:"required" label:"标签名称"`
	TagStatus int    `json:"tagStatus" validate:"required,oneof=10 20" label:"标签状态"`
}

type TagUpdateModel struct {
	Id int `validate:"required,gt=0"`
	TagName   string `validate:"required,max=32" label:"标签名称"`
	TagStatus int    `validate:"required,oneof=10 20" label:"标签状态"`
}

type TagDeleteModel struct {
	Id int `validate:"required,gt=0"`
}

type TagListForm struct {
	Page int `validate:"omitempty,gte=1" label:"当前页码"`
	PageSize int `validate:"omitempty,gte=1" label:"每页条数"`
	TagName string `validate:"omitempty,max=32" label:"标签名称"`
}

