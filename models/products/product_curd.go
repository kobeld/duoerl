package products

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
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

func FindByBrandId(brandId bson.ObjectId) (r []*Product, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	query := bson.M{"brandid": brandId}
	err = mgodb.FindAll(PRODUCTS, query, &r)
	return

}

func FindById(id bson.ObjectId) (product *Product, err error) {
	if !id.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindByIds(ids []bson.ObjectId) (products []*Product, err error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindOne(query bson.M) (r *Product, err error) {
	err = mgodb.FindOne(PRODUCTS, query, &r)
	return
}

func FindAll(query bson.M) (r []*Product, err error) {
	err = mgodb.FindAll(PRODUCTS, query, &r)
	return
}

func CountProductByBrandId(brandId bson.ObjectId) (num int, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
	}
	return CountProduct(bson.M{"brandid": brandId})
}

func CountProduct(query bson.M) (num int, err error) {
	mgodb.CollectionDo(PRODUCTS, func(c *mgo.Collection) {
		num, err = c.Find(query).Count()
	})
	return
}
