package products

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
)

const PRODUCTS = "products"

func (this *Product) Save() error {
	return mgodb.Save(PRODUCTS, this)
}

func FindById(id bson.ObjectId) (*Product, error) {
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
