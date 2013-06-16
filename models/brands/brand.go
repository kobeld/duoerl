package brands

import (
	"labix.org/v2/mgo/bson"
)

type Brand struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string
	Alias   string // Another name that may be different language
	Intro   string
	Country string
	Website string
	LogoUrl string
}

func (this *Brand) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

// func NewBrandFromInput(brandInput *duoerlapi.BrandInput) *Brand {
// 	return Brand{
// 		Id
// 	}
// }
