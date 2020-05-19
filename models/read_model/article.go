package read_model

import "go-gin-blog-api/models"

type Article struct {
	Id            int           `json:"id" gorm:"column:id;primary_key"`
	Title         string        `json:"title" gorm:"column:title"`
	Content       string        `json:"content" gorm:"column:content"`
	CurrentStatus int           `json:"currentStatus" gorm:"column:current_status"`
	Tags          []Tag  `json:"tags" gorm:"many2many:article_tag"`
	ArticleAttachments []models.ArticleAttachment `json:"articleAttachments" gorm:"foreignkey:ArticleId;AssociationForeignKey:Id"`
}

// 标签
type Tag struct {
	Id        int           `json:"id" gorm:"column:id;primary_key"`
	TagName   string        `json:"tagName" gorm:"column:tag_name"`
	TagStatus int           `json:"tagStatus" gorm:"column:tag_status"`
}

func (article *Article) Set(tagIds []int) *Article {
	var tags []Tag
	for _, tagId := range tagIds {
		tags = append(tags, Tag{Id: tagId})
	}
	article.Tags = tags
	return article
}
