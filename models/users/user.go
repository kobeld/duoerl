package users

import (
	"fmt"
	"github.com/kobeld/duoerl/global"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id              bson.ObjectId `bson:"_id"`
	Name            string
	Email           string
	Password        string
	ConfirmPassword string `bson:"-" json:"-"`
	WishProductIds  []bson.ObjectId
	OwnProductIds   []bson.ObjectId
	FollowBrandIds  []bson.ObjectId
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

// --------- Methods ----------

func (this *User) MakeId() interface{} {
	if this.Id == "" {
		this.Id = bson.NewObjectId()
	}
	return this.Id
}

func (this *User) Link() string {
	return fmt.Sprintf("/user/%s", this.Id.Hex())
}

func (this *Profile) BirthdayText() string {
	if this.Birthday.IsZero() {
		return global.TEXT_BIRTHDAY_SECRET
	}
	return this.Birthday.Format(global.DATE_BIRTHDAY)
}

func (this *Profile) GenderText() string {
	if this.Gender {
		return global.TEXT_GENDER_MALE
	}

	return global.TEXT_GENDER_FEMALE
}

func (this *User) HasWishedProduct(productId bson.ObjectId) bool {
	for _, id := range this.WishProductIds {
		if id == productId {
			return true
		}
	}
	return false
}

func (this *User) HasFollowedBrand(brandId bson.ObjectId) bool {
	for _, id := range this.FollowBrandIds {
		if id == brandId {
			return true
		}
	}
	return false
}

// --------- Functions ----------

func NewUser() *User {
	return &User{}
}

func BuildUserMap(dbUsers []*User) map[bson.ObjectId]*User {
	userMap := make(map[bson.ObjectId]*User)
	for _, dbUser := range dbUsers {
		userMap[dbUser.Id] = dbUser
	}

	return userMap
}
