package services

import (
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/models/wishitems"
	"github.com/kobeld/duoerl/utils"
)

func CreateWishItem(userId, productId string) (err error) {

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

	if err = users.AddWishProduct(userOId, wishItem.ProductId); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func DeleteWishItem(userId, productId string) (err error) {

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

	err = users.RemoveWishProduct(userOId, productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetWishItem(userId, productId string) (wishItem *wishitems.WishItem) {

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

	wishItem, _ = wishitems.FindByUserAndProductId(userOId, productOId)

	return

}
