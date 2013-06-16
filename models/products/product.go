package products

import (
	"labix.org/v2/mgo/bson"
)

type Product struct {
	Id      bson.ObjectId `bson:"_id"`
	BrandId bson.ObjectId
	Name    string
	Alias   string
	Intro   string
	Image   string
}

func (this *Product) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}
