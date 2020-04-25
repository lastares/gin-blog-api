package http

type Failed struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
}