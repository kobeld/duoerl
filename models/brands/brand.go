package brands

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"time"
)

type Brand struct {
	Id        bson.ObjectId `bson:"_id"`
	Name      string
	Alias     string // Another name that may be different language
	Intro     string
	Country   string
	Website   string
	Logo      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *Brand) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *Brand) Link() string {
	return fmt.Sprintf("/brand/%s", this.Id.Hex())
}

func (this *Brand) LogoUrl() string {
	if this.Logo == "" {
		return "http://lorempixel.com/g/200/200/"
	}
	return this.Logo
}

func BuildBrandMap(dbBrands []*Brand) map[bson.ObjectId]*Brand {
	brandMap := make(map[bson.ObjectId]*Brand)
	for _, dbBrand := range dbBrands {
		brandMap[dbBrand.Id] = dbBrand
	}
	return brandMap
}
