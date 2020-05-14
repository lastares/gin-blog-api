package response

type Succeed struct {
	Code int `json:"code" example:"200"`
}

type ResponseJson struct {
	Code int `json:"code" example:"0"`
	Data interface{} `json:"data"`
	Meta Meta `json:"meta"`
}

type Failed struct {
	Code int    `json:"code" example:"-1"`
	Msg  string `json:"msg" example:"failed"`
}

type Meta struct {
	Total int `json:"total"`
}
