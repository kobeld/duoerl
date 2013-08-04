package users

import (
	"github.com/kobeld/duoerl/utils"
	"github.com/sunfmin/mgodb"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
	"time"
)

const USERS = "users"

func (this *User) Save() error {
	if this.CreatedAt.IsZero() {
		this.CreatedAt = time.Now()
	} else {
		this.UpdatedAt = time.Now()
	}
	this.Email = strings.ToLower(this.Email)
	return mgodb.Save(USERS, this)
}

func FetchByIdHex(idHex string) (user *User, err error) {
	userId, err := utils.ToObjectId(idHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	user, err = FindById(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func FindById(id bson.ObjectId) (user *User, err error) {
	if !id.Valid() {
		return
	}
	return FindOne(bson.M{"_id": id})
}

func FindByIds(ids []bson.ObjectId) ([]*User, error) {
	return FindAll(bson.M{"_id": bson.M{"$in": ids}})
}

func FindByEmail(email string) (*User, error) {
	return FindOne(bson.M{"email": strings.ToLower(email)})
}

func FindOne(query bson.M) (user *User, err error) {
	err = mgodb.FindOne(USERS, query, &user)
	return
}

func FindAll(query bson.M) (users []*User, err error) {
	err = mgodb.FindAll(USERS, query, &users)
	return
}

func Update(selector, changer bson.M) (err error) {
	mgodb.CollectionDo(USERS, func(rc *mgo.Collection) {
		err = rc.Update(selector, changer)
	})
	return
}

// ----- Duplicated, just for reference
func AddOwnProduct(userId, productId bson.ObjectId) (err error) {
	selector := bson.M{"_id": userId}
	changer := bson.M{"$push": bson.M{"ownproductids": productId}}
	return Update(selector, changer)
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

func AddFollowBrand(userId, brandId bson.ObjectId) (err error) {
	selector := bson.M{"_id": userId}
	changer := bson.M{"$push": bson.M{"followbrandids": brandId}}
	return Update(selector, changer)
}

func RemoveFollowBrand(userId, brandId bson.ObjectId) (err error) {
	selector := bson.M{"_id": userId}
	changer := bson.M{"$pull": bson.M{"followbrandids": brandId}}
	return Update(selector, changer)
}
