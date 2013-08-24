package notes

import (
	"fmt"
	"github.com/kobeld/duoerl/models/articles"
	"labix.org/v2/mgo/bson"
)

type Note struct {
	Id               bson.ObjectId `bson:"_id"`
	articles.Article `,inline`
}

func (this *Note) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *Note) Link() string {
	return fmt.Sprintf("/note/%s", this.Id.Hex())
}
