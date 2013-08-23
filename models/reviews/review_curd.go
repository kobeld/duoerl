package reviews

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
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
		err = global.InvalidIdError
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindSomeByBrandId(brandId bson.ObjectId) (rs []*Review, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}

	return FindAll(bson.M{"brandid": brandId})
}

func FindSomeByProductId(productId bson.ObjectId) (rs []*Review, err error) {
	if !productId.Valid() {
		err = global.InvalidIdError
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

func CountReviewByProductId(productId bson.ObjectId) (num int, err error) {
	if !productId.Valid() {
		err = global.InvalidIdError
		return
	}
	return CountReview(bson.M{"productid": productId})
}

func CountReviewByBrandId(brandId bson.ObjectId) (num int, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return CountReview(bson.M{"brandid": brandId})
}

func CountReview(query bson.M) (num int, err error) {
	mgodb.CollectionDo(REVIEWS, func(c *mgo.Collection) {
		num, err = c.Find(query).Count()
	})
	return
}
