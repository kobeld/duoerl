package ownitems

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"time"
)

const OWN_ITEMS = "own_items"

func (this *OwnItem) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	}
	return mgodb.Save(OWN_ITEMS, this)
}

func FindByUserAndProductId(userId, productId bson.ObjectId) (*OwnItem, error) {
	return FindOne(bson.M{"userid": userId, "productid": productId})
}

func FindOne(query bson.M) (r *OwnItem, err error) {
	err = mgodb.FindOne(OWN_ITEMS, query, &r)
	return
}

func DeleteByUserAndProductId(userId, productId bson.ObjectId) error {
	return DeleteOwnItem(bson.M{"userid": userId, "productid": productId})
}

func DeleteOwnItem(query bson.M) (err error) {
	mgodb.CollectionDo(OWN_ITEMS, func(rc *mgo.Collection) {
		_, err = rc.RemoveAll(query)
	})
	return
}
