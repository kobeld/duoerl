package users

import (
	"fmt"
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/utils/gravatar"
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id              bson.ObjectId `bson:"_id"`
	Name            string
	Email           string
	Password        string
	ConfirmPassword string `bson:"-" json:"-"`
	AvatarUrl       string
	Profile         Profile
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Profile struct {
	Gender      bool
	Description string
	Location    string
	Birthday    time.Time
	SkinTexture string
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

func (this *Profile) GenderText() string {
	if this.Gender {
		return global.TEXT_GENDER_MALE
	}

	return global.TEXT_GENDER_FEMALE
}

func (this *User) Avatar() string {
	if this.AvatarUrl != "" {
		return this.AvatarUrl
	}
	return gravatar.UrlSize(this.Email, 160)
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
