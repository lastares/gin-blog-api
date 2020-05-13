package models

type Comment struct {
	CommentContent string `json:"commentContent"`
	ArticleId int `json:"articleId"`
	UserId int `json:"userId"`
}
