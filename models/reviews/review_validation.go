package reviews

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/utils"
	"github.com/sunfmin/govalidations"
	"labix.org/v2/mgo/bson"
)

func (this *Review) ValidateLikeAction(userId bson.ObjectId) *govalidations.Validated {
	gk := &ReviewGateKeeper{}
	gk.AddLikeValidator(userId)
	return gk.Validate(this)
}

type ReviewGateKeeper struct {
	govalidations.GateKeeper
}

func (this *ReviewGateKeeper) AddLikeValidator(userId bson.ObjectId) {
	this.Add(govalidations.Custom(func(object interface{}) bool {
		return !utils.IsInObjectIds(userId, object.(*Review).LikedByIds)
	}, "LikedByIds", global.REVIEW_01))
}
