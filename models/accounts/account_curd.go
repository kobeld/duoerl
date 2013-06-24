package accounts

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

const ACCOUNTS = "accounts"

func (this *Account) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	this.Email = strings.ToLower(this.Email)
	return mgodb.Save(ACCOUNTS, this)
}

func AddWishProduct(userId, productId bson.ObjectId) (err error) {
	selector := bson.M{"_id": userId}
	changer := bson.M{"$push": bson.M{"wishproductids": productId}}
	return Update(selector, changer)
}

func RemoveWishProduct(userId, productId bson.ObjectId) (err error) {
	selector := bson.M{"_id": userId}
	changer := bson.M{"$pull": bson.M{"wishproductids": productId}}
	return Update(selector, changer)
}

func FindById(id bson.ObjectId) (account *Account, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindByIds(ids []bson.ObjectId) ([]*Account, error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindByEmail(email string) (*Account, error) {
	return FindOne(bson.M{"email": strings.ToLower(email)})
}

func FindOne(query bson.M) (account *Account, err error) {
	err = mgodb.FindOne(ACCOUNTS, query, &account)
	return
}

func FindAll(query bson.M) (accounts []*Account, err error) {
	err = mgodb.FindAll(ACCOUNTS, query, &accounts)
	return
}

func Update(selector, changer bson.M) (err error) {
	mgodb.CollectionDo(ACCOUNTS, func(rc *mgo.Collection) {
		err = rc.Update(selector, changer)
	})
	return
}
