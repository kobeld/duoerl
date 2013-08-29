package posts

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Post struct {
	Id        bson.ObjectId `bson:"_id"`
	Content   string
	AuthorId  bson.ObjectId
	CreatedAt time.Time
}

func (this *Post) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
