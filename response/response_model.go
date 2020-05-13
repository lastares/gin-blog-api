package response

type Succeed struct {
	Code int `json:"code" example:"200"`
}

type ResponseJson struct {
	Code int `json:"code" example:"0"`
	Data interface{} `json:"data"`
}

type Failed struct {
	Code int    `json:"code" example:"-1"`
	Msg  string `json:"msg" example:"failed"`
}
