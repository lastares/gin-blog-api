package dto

type TagList struct {
	Code int `json:"code" example:"200"`
	Data interface{} `json:"data" example:"Tag"`
}
