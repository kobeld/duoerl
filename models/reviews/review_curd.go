package reviews

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const REVIEWS = "reviews"

func (this *Review) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	}
	return mgodb.Save(REVIEWS, this)
}

func FindById(id bson.ObjectId) (review *Review, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindSomeByProductId(productId bson.ObjectId) (rs []*Review, err error) {
	if !productId.Valid() {
		return
	}

	return FindAll(bson.M{"productid": productId})
}

func FindOne(query bson.M) (review *Review, err error) {
	err = mgodb.FindOne(REVIEWS, query, &review)
	return
}

func FindAll(query bson.M) (review []*Review, err error) {
	err = mgodb.FindAll(REVIEWS, query, &review)
	return
}
