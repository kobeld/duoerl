package products

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const PRODUCTS = "products"

func (this *Product) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	return mgodb.Save(PRODUCTS, this)
}

func FindById(id bson.ObjectId) (product *Product, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindOne(query bson.M) (r *Product, err error) {
	err = mgodb.FindOne(PRODUCTS, query, &r)
	return
}

func FindAll(query bson.M) (r []*Product, err error) {
	err = mgodb.FindAll(PRODUCTS, query, &r)
	return
}
