package wishitems

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"time"
)

const WISH_ITEMS = "wish_items"

func (this *WishItem) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	}
	return mgodb.Save(WISH_ITEMS, this)
}

func FindById(id bson.ObjectId) (wishItem *WishItem, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindByUserId(userId bson.ObjectId) (wishItems []*WishItem, err error) {
	if !userId.Valid() {
		return
	}
	return FindAll(bson.M{"userid": userId})
}

func FindByIds(ids []bson.ObjectId) (wishItems []*WishItem, err error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindOne(query bson.M) (r *WishItem, err error) {
	err = mgodb.FindOne(WISH_ITEMS, query, &r)
	return
}

func FindAll(query bson.M) (r []*WishItem, err error) {
	err = mgodb.FindAll(WISH_ITEMS, query, &r)
	return
}
