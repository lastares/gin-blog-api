package http

type Success struct {
	Code int `json:"code"`
	Data interface{}
}
