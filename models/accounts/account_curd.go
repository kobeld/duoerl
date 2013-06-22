package accounts

import (
	"github.com/sunfmin/mgodb"
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
