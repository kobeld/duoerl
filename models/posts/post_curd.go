package posts

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const (
	POSTS = "posts"
)

func (this *Post) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	}
	return mgodb.Save(POSTS, this)
}

func FindSomeByAuthorId(authorId bson.ObjectId) (r []*Post, err error) {
	if !authorId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"authorid": authorId})
}

func FindOne(query bson.M) (post *Post, err error) {
	err = mgodb.FindOne(POSTS, query, &post)
	return
}

func FindAll(query bson.M) (r []*Post, err error) {
	err = mgodb.FindAll(POSTS, query, &r)
	return
}
