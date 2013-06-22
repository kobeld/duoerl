package products

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Product struct {
	Id        bson.ObjectId `bson:"_id"`
	BrandId   bson.ObjectId
	AuthorId  bson.ObjectId
	Name      string
	Alias     string
	Intro     string
	Image     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (this *Product) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func CollectBrandAndAuthorIds(dbProducts []*Product) (brandIds, authorIds []bson.ObjectId) {
	for _, dbProduct := range dbProducts {
		if dbProduct.BrandId.Valid() {
			brandIds = append(brandIds, dbProduct.BrandId)
		}
		if dbProduct.AuthorId.Valid() {
			authorIds = append(authorIds, dbProduct.AuthorId)
		}
	}
	return
}
