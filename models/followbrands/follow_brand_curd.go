package followbrands

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const FOLLOW_BRANDS = "follow_brands"

func (this *FollowBrand) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	}
	return mgodb.Save(FOLLOW_BRANDS, this)
}

func FindByBrandId(brandId bson.ObjectId) (r []*FollowBrand, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"brandid": brandId})
}

func FindByUserAndBrandId(userId, brandId bson.ObjectId) (followBrand *FollowBrand, err error) {
	if !userId.Valid() || !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindOne(bson.M{"userid": userId, "brandid": brandId})
}

func FindByUserId(userId bson.ObjectId) (r []*FollowBrand, err error) {
	if !userId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"userid": userId})
}

func FindByIds(ids []bson.ObjectId) ([]*FollowBrand, error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindOne(query bson.M) (r *FollowBrand, err error) {
	err = mgodb.FindOne(FOLLOW_BRANDS, query, &r)
	return
}

func FindAll(query bson.M) (r []*FollowBrand, err error) {
	err = mgodb.FindAll(FOLLOW_BRANDS, query, &r)
	return
}

func DeleteFollowBrand(query bson.M) (err error) {
	mgodb.CollectionDo(FOLLOW_BRANDS, func(rc *mgo.Collection) {
		_, err = rc.RemoveAll(query)
	})
	return
}

func CountBrandFollowerByBrandId(brandId bson.ObjectId) (num int, err error) {
	if !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return CountFollowBrand(bson.M{"brandid": brandId})
}

func CountFollowingBrandByUserId(userId bson.ObjectId) (num int, err error) {
	if !userId.Valid() {
		err = global.InvalidIdError
		return
	}
	return CountFollowBrand(bson.M{"userid": userId})
}

func CountFollowBrand(query bson.M) (num int, err error) {
	mgodb.CollectionDo(FOLLOW_BRANDS, func(c *mgo.Collection) {
		num, err = c.Find(query).Count()
	})
	return
}

func DeleteByUserAndBrandId(userId, brandId bson.ObjectId) (err error) {
	if !userId.Valid() || !brandId.Valid() {
		err = global.InvalidIdError
		return
	}
	return DeleteFollowBrand(bson.M{"userid": userId, "brandid": brandId})
}
