package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/categories"
	"github.com/kobeld/duoerl/models/efficacies"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
)

func GetFullCategories() []*duoerlapi.Category {

	// Return the cached categories
	if len(global.Categories) == 0 {
		global.Categories = getCategories()
	}

	return global.Categories
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

func resetCategories() {
	global.Categories = getCategories()
}

func getCategories() (apiCategories []*duoerlapi.Category) {

	apiCategoryMap := make(map[string]*duoerlapi.Category)
	apiSubCategories := []*duoerlapi.SubCategory{}

	// Get all Categories
	allCategories, err := categories.FindAll(nil)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, category := range allCategories {
		switch category.Level {
		case categories.LEVEL_ONE:
			apiCategory := toApiCategory(category)
			apiCategoryMap[apiCategory.Id] = apiCategory
			apiCategories = append(apiCategories, apiCategory)

		case categories.LEVEL_TWO:
			apiSubCategories = append(apiSubCategories, toApiSubCategory(category))
		}
	}

	for _, apiSubCategory := range apiSubCategories {
		if apiCategory, exist := apiCategoryMap[apiSubCategory.ParentId]; exist {
			apiCategory.SubCategories = append(apiCategory.SubCategories, apiSubCategory)
		}
	}

	// Get all Efficacies
	allEfficacies, err := efficacies.FindAll(nil)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, efficacy := range allEfficacies {
		apiEfficacy := toApiEfficacy(efficacy)
		if apiCategory, exist := apiCategoryMap[apiEfficacy.ParentId]; exist {
			apiCategory.Efficacies = append(apiCategory.Efficacies, apiEfficacy)
		}
	}

	return
}

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
