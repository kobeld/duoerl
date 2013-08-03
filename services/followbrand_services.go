package services

import (
	"github.com/kobeld/duoerl/models/followbrands"
	"github.com/kobeld/duoerl/utils"
)

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
