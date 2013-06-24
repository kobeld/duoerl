package followbrands

import (
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

func DeleteByUserAndBrandId(userId, brandId bson.ObjectId) error {
	return DeleteFollowBrand(bson.M{"userid": userId, "brandid": brandId})
}

func FindByUserAndBrandId(userId, brandId bson.ObjectId) (*FollowBrand, error) {
	return FindOne(bson.M{"userid": userId, "brandid": brandId})
}

func FindByUserId(userId bson.ObjectId) (r []*FollowBrand, err error) {
	if !userId.Valid() {
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
