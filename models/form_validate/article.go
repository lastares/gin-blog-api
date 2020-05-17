package form_validate

// 文章校验结构体
type ArticleCreateForm struct {
	Title   string `validate:"required,max=256" label:"文章标题"`
	Content string    `validate:"required" label:"文章内容"`
	CurrentStatus int    `validate:"required,oneof=10 20 30" label:"文章状态"`
	TagIds []int    `validate:"required" label:"标签"`
	Attachments []string `validate:"required,max=256" label:"附件地址"`

}

type ArticleUpdateForm struct {
	Id int `validate:"required,gte=1" label:"ID"`
	Title   string `validate:"required,max=256" label:"文章标题"`
	Content string    `validate:"required" label:"文章内容"`
	CurrentStatus int    `validate:"required,oneof=10 20 30" label:"文章状态"`
	TagIds []int    `validate:"required" label:"标签"`
	Attachments []string `validate:"required,max=256" label:"附件地址"`
}

type ArticleListForm struct {
	Page int `validate:"omitempty,gte=1" label:"当前页码"`
	PageSize int `validate:"omitempty,gte=1" label:"每页条数"`
	Content string `validate:"omitempty,max=32" label:"文章内容"`
}

type ArticleDeleteForm struct {
	Id int `validate:"required,gte=1" label:"ID"`

}



