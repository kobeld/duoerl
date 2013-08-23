package services

import (
	"github.com/jmcvetta/randutil"
	"github.com/kobeld/duoerl/configs"
	"github.com/kobeld/duoerl/models/followbrands"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

func GetBrandFollowers(brandIdHex string) (apiUsers []*duoerlapi.User, err error) {

	brandId, err := utils.ToObjectId(brandIdHex)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	followbrandz, err := followbrands.FindByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	maxNum := len(followbrandz)
	// Get random number users
	if maxNum > configs.BRAND_SHOW_FOLLOWER_NUM {
		randIndex, err := randutil.IntRange(0, maxNum)
		if err != nil {
			utils.PrintStackAndError(err)
			randIndex = 0
		}

		leftIndex := randIndex - configs.BRAND_SHOW_FOLLOWER_NUM
		if leftIndex < 0 {
			followbrandz = followbrandz[0:configs.BRAND_SHOW_FOLLOWER_NUM]
		} else {
			followbrandz = followbrandz[leftIndex:randIndex]
		}
	}

	followerIds := []bson.ObjectId{}
	for _, followBrand := range followbrandz {
		followerIds = append(followerIds, followBrand.UserId)
	}

	followers, err := users.FindByIds(followerIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiUsers = toApiUsers(followers)

	return
}

func CreateFollowBrand(userId, brandId string) (err error) {

	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Validation, check if the record has been created
	followBrand, err := followbrands.FindByUserAndBrandId(userOId, brandOId)
	if followBrand != nil {
		return
	}

	followBrand = &followbrands.FollowBrand{
		UserId:  userOId,
		BrandId: brandOId,
	}

	if err = followBrand.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func DeleteFollowBrand(userId, brandId string) (err error) {

	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	err = followbrands.DeleteByUserAndBrandId(userOId, brandOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetFollowBrand(userId, brandId string) (followBrand *followbrands.FollowBrand) {
	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	followBrand, _ = followbrands.FindByUserAndBrandId(userOId, brandOId)

	return
}
