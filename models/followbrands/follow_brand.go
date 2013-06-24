package followbrands

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type FollowBrand struct {
	Id        bson.ObjectId `bson:"_id"`
	UserId    bson.ObjectId
	BrandId   bson.ObjectId
	CreatedAt time.Time
}

func (this *FollowBrand) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
