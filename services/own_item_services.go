package services

import (
	"github.com/kobeld/duoerl/models/ownitems"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
)

func AddOwnItem(ownItemInput *duoerlapi.OwnItemInput) (err error) {

	userOId, err := utils.ToObjectId(ownItemInput.UserId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productOId, err := utils.ToObjectId(ownItemInput.ProductId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Validation, check if the record has been created
	ownItem, err := ownitems.FindByUserAndProductId(userOId, productOId)
	if ownItem != nil {
		return
	}

	ownItem = &ownitems.OwnItem{
		UserId:    userOId,
		ProductId: productOId,
		GotFrom:   ownItemInput.GotFrom,
	}

	if err = ownItem.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func RemoveOwnItem(userId, productId string) (err error) {
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

	err = ownitems.DeleteByUserAndProductId(userOId, productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func GetOwnItem(userId, productId string) (ownItem *ownitems.OwnItem, err error) {

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

	ownItem, _ = ownitems.FindByUserAndProductId(userOId, productOId)

	return

}
