package accounts

import (
	// "github.com/sunfmin/mgodb"
	// "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Account struct {
	Id              bson.ObjectId `bson:"_id"`
	Name            string
	Email           string
	Password        string
	ConfirmPassword string `bson:"-"`
}

func (this *Account) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func NewAccount() *Account {
	return &Account{}
}
