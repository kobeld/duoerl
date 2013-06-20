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

func CollectBrandIds(dbProducts []*Product) (brandIds []bson.ObjectId) {
	for _, dbProduct := range dbProducts {
		brandIds = append(brandIds, dbProduct.BrandId)
	}
	return
}
