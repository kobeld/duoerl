package articles

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Article struct {
	Title     string
	Content   string
	AuthorId  bson.ObjectId
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewArticle(title, content string, authorId bson.ObjectId) *Article {
	return &Article{
		Title:    title,
		Content:  content,
		AuthorId: authorId,
	}
}
