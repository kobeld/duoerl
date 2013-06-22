package brands

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const BRANDS = "brands"

func (this *Brand) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	return mgodb.Save(BRANDS, this)
}

func FindById(id bson.ObjectId) (brand *Brand, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindByIds(ids []bson.ObjectId) (brands []*Brand, err error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindOne(query bson.M) (brand *Brand, err error) {
	err = mgodb.FindOne(BRANDS, query, &brand)
	return
}

func FindAll(query bson.M) (brands []*Brand, err error) {
	err = mgodb.FindAll(BRANDS, query, &brands)
	return
}
