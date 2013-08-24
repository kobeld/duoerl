package news

import (
	"fmt"
	"github.com/kobeld/duoerl/models/articles"
	"labix.org/v2/mgo/bson"
)

type News struct {
	Id               bson.ObjectId `bson:"_id"`
	BrandId          bson.ObjectId
	articles.Article `,inline`
}

func (this *News) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *News) Link() string {
	return fmt.Sprintf("/news/%s", this.Id.Hex())
}
