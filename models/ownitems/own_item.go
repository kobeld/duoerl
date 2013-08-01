package ownitems

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type OwnItem struct {
	Id        bson.ObjectId `bson:"_id"`
	UserId    bson.ObjectId
	ProductId bson.ObjectId
	CreatedAt time.Time

	GotFrom string
}

func (this *OwnItem) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
