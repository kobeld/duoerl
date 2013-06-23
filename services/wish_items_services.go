package services

import (
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

	wishItem := &wishitems.WishItem{
		UserId:    userOId,
		ProductId: productOId,
	}

	if err = wishItem.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}
