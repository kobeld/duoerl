package accounts

import (
	"github.com/kobeld/duoerl/global"
	"labix.org/v2/mgo/bson"
	"time"
)

type Account struct {
	Id              bson.ObjectId `bson:"_id"`
	Name            string
	Email           string
	Password        string
	ConfirmPassword string `bson:"-" json:"-"`
	WishProductIds  []bson.ObjectId
	Profile         Profile
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Profile struct {
	Gender      bool
	Description string
	Location    string
	Birthday    time.Time
	HairTexture string
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

func BuildAccountMap(dbAccounts []*Account) map[bson.ObjectId]*Account {
	accountMap := make(map[bson.ObjectId]*Account)
	for _, dbAccount := range dbAccounts {
		accountMap[dbAccount.Id] = dbAccount
	}

	return accountMap
}

func (this *Account) Birthday() string {
	if this.Profile.Birthday.IsZero() {
		return global.TEXT_BIRTHDAY_SECRET
	}
	return this.Profile.Birthday.Format(global.DATE_BIRTHDAY)
}

func (this *Account) Gender() string {
	if !this.Profile.Gender {
		return global.TEXT_GENDER_FEMALE
	}
	return global.TEXT_GENDER_MALE
}

func (this *Account) HasWishedProduct(productId bson.ObjectId) bool {
	for _, id := range this.WishProductIds {
		if id == productId {
			return true
		}
	}
	return false
}
