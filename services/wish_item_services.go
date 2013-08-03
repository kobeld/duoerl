package services

import (
	"github.com/kobeld/duoerl/models/wishitems"
	"github.com/kobeld/duoerl/utils"
)

func AddWishItem(userId, productId string) (err error) {

	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productOId, err := utils.ToObjectId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Validation, check if the record has been created
	wishItem, err := wishitems.FindByUserAndProductId(userOId, productOId)
	if wishItem != nil {
		return
	}

	wishItem = &wishitems.WishItem{
		UserId:    userOId,
		ProductId: productOId,
	}

	if err = wishItem.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func RemoveWishItem(userId, productId string) (err error) {

	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productOId, err := utils.ToObjectId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	err = wishitems.DeleteByUserAndProductId(userOId, productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetWishItem(userId, productId string) (wishItem *wishitems.WishItem, err error) {

	userOId, err := utils.ToObjectId(userId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productOId, err := utils.ToObjectId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	wishItem, err = wishitems.FindByUserAndProductId(userOId, productOId)

	return

}
