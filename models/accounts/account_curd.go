package accounts

import (
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo/bson"
	"strings"
)

const ACCOUNTS = "accounts"

func (this *Account) Save() error {
	this.Email = strings.ToLower(this.Email)
	return mgodb.Save(ACCOUNTS, this)
}

func FindByEmail(email string) (*Account, error) {
	return FindOne(bson.M{"email": strings.ToLower(email)})
}

func FindOne(query bson.M) (account *Account, err error) {
	err = mgodb.FindOne(ACCOUNTS, query, &account)
	return
}
