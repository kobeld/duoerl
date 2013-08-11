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

func GetFullCategories() (r []*duoerlapi.Category, err error) {

	allCategories, err := categories.FindAll(nil)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiCategoryMap := make(map[string]*duoerlapi.Category)
	apiSubCategories := []*duoerlapi.SubCategory{}

	for _, category := range allCategories {
		switch category.Level {
		case categories.LEVEL_ONE:
			apiCategory := toApiCategory(category)
			apiCategoryMap[apiCategory.Id] = apiCategory
			r = append(r, apiCategory)

		case categories.LEVEL_TWO:
			apiSubCategories = append(apiSubCategories, toApiSubCategory(category))
		}
	}

	for _, apiSubCategory := range apiSubCategories {
		if apiCategory, exist := apiCategoryMap[apiSubCategory.ParentId]; exist {
			apiCategory.SubCategories = append(apiCategory.SubCategories, apiSubCategory)
		}
	}

	return
}

// Not using yet
func GetClassifiedCategories() (r *duoerlapi.ClassifiedCategories, err error) {
	r = new(duoerlapi.ClassifiedCategories)

	allCategories, err := categories.FindAll(nil)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, category := range allCategories {
		switch category.Level {
		case categories.LEVEL_ONE:
			r.Categories = append(r.Categories, toApiCategory(category))
		case categories.LEVEL_TWO:
			r.SubCategories = append(r.SubCategories, toApiSubCategory(category))
		}
	}

	return
}

// -------------

func toApiCategory(category *categories.Category) *duoerlapi.Category {
	apiCategory := new(duoerlapi.Category)
	if category != nil {
		apiCategory = &duoerlapi.Category{
			Id:   category.Id.Hex(),
			Name: category.Name,
		}
	}
	return apiCategory
}

func toApiSubCategory(category *categories.Category) *duoerlapi.SubCategory {
	apiSubCategory := new(duoerlapi.SubCategory)
	if category != nil {
		apiSubCategory = &duoerlapi.SubCategory{
			Id:       category.Id.Hex(),
			Name:     category.Name,
			ParentId: category.ParentId.Hex(),
		}
	}
	return apiSubCategory
}
