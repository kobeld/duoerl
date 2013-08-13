package services

import (
	"github.com/kobeld/duoerl/models/categories"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
)

func CreateCategory(input *duoerlapi.CategoryInput) (originInput *duoerlapi.CategoryInput, err error) {

	originInput = input
	category := categories.NewCategory(input)

	if validated := category.ValidateCreation(); validated.HasError() {
		err = validated.ToError()
		return
	}

	if err = category.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}
