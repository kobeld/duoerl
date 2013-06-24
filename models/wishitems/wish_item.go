package wishitems

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type WishItem struct {
	Id        bson.ObjectId `bson:"_id"`
	UserId    bson.ObjectId
	ProductId bson.ObjectId
	CreatedAt time.Time
}

func (this *WishItem) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
