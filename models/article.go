package models

import "go-gin-blog-api/util"

const (
	// 文章状态：正常
	ARTICLE_STATUS_NORMAL = 10

	// 文章状态：草稿
	ARTICLE_STATUS_DRAFT = 20

	// 文章状态：废弃
	ARTICLE_STATUS_DROP = 30
)


// 文章
type Article struct {
	Id                 int                 `json:"id" gorm:"column:id;primary_key"`
	Title              string              `json:"title" gorm:"column:title"`
	Content            string              `json:"content" gorm:"column:content"`
	CurrentStatus      int                 `json:"currentStatus" gorm:"column:current_status"`
	Tags               []Tag               `json:"tags" gorm:"many2many:article_tag;association_autoupdate:false;"`
	ArticleAttachments []ArticleAttachment `json:"articleAttachments" gorm:"foreignkey:ArticleId;AssociationForeignKey:Id"`
	CreatedAt          util.JSONTime       `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt          util.JSONTime       `json:"updatedAt" gorm:"column:updated_at"`
}

func (article *Article) SetTitle(title string) *Article {
	article.Title = title
	return article
}

func (aricle *Article) SetContent(content string) *Article {
	aricle.Content = content
	return aricle
}

func (article *Article) SetCurrentStatus(currentStatus int) *Article {
	article.CurrentStatus = currentStatus
	return article
}

func (article *Article) SetTags(tagIds []int) *Article {
	var tags []Tag
	for _, tagId := range tagIds {
		tags = append(tags, Tag{Id: tagId})
	}
	article.Tags = tags
	return article
}

func (article *Article) SetAttachments(attachmentPaths []string) *Article {
	var articleAttachments []ArticleAttachment
	for _, attachmentPath := range attachmentPaths {
		articleAttachments = append(articleAttachments, ArticleAttachment{AttachmentPath: attachmentPath})
	}
	article.ArticleAttachments = articleAttachments
	return article
}

func (article *Article) GetTags() []Tag {
	return article.Tags
}

func (article *Article) GetAttachments() []ArticleAttachment {
	return article.ArticleAttachments
}
