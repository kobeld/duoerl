package wishitems

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
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

func DeleteByUserAndProductId(userId, productId bson.ObjectId) error {
	return DeleteWishItem(bson.M{"userid": userId, "productid": productId})
}

func FindByUserAndProductId(userId, productId bson.ObjectId) (*WishItem, error) {
	return FindOne(bson.M{"userid": userId, "productid": productId})
}

func FindByUserId(userId bson.ObjectId) (r []*WishItem, err error) {
	if !userId.Valid() {
		err = global.InvalidIdError
		return
	}
	return FindAll(bson.M{"userid": userId})
}

func FindByIds(ids []bson.ObjectId) ([]*WishItem, error) {
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

func DeleteWishItem(query bson.M) (err error) {
	mgodb.CollectionDo(WISH_ITEMS, func(rc *mgo.Collection) {
		_, err = rc.RemoveAll(query)
	})
	return
}
