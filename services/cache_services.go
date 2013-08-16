package services

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/models/categories"
	"github.com/kobeld/duoerl/models/efficacies"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
)

func GetCategories() []*duoerlapi.Category {
	if len(global.Categories) == 0 {
		populateCachedCategoriesRelated()
	}
	return global.Categories
}

func GetCategoryMap() map[string]*duoerlapi.Category {
	if len(global.CategoryMap) == 0 {
		populateCachedCategoriesRelated()
	}
	return global.CategoryMap
}

func GetSubCategoryMap() map[string]*duoerlapi.SubCategory {
	if len(global.SubCategoryMap) == 0 {
		populateCachedCategoriesRelated()
	}
	return global.SubCategoryMap
}

func GetEfficacyMap() map[string]*duoerlapi.Efficacy {
	if len(global.EfficacyMap) == 0 {
		populateCachedCategoriesRelated()
	}
	return global.EfficacyMap
}

// -------------

func resetCategories() {
	populateCachedCategoriesRelated()
}

func populateCachedCategoriesRelated() {

	apiCategoryMap := make(map[string]*duoerlapi.Category)
	apiSubCategoryMap := make(map[string]*duoerlapi.SubCategory)
	apiEfficacyMap := make(map[string]*duoerlapi.Efficacy)

	apiCategories := []*duoerlapi.Category{}
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
			apiSubCategory := toApiSubCategory(category)
			apiSubCategories = append(apiSubCategories, apiSubCategory)
			apiSubCategoryMap[apiSubCategory.Id] = apiSubCategory
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
		apiEfficacyMap[apiEfficacy.Id] = apiEfficacy
		if apiCategory, exist := apiCategoryMap[apiEfficacy.ParentId]; exist {
			apiCategory.Efficacies = append(apiCategory.Efficacies, apiEfficacy)
		}
	}

	global.Categories = apiCategories
	global.CategoryMap = apiCategoryMap
	global.SubCategoryMap = apiSubCategoryMap
	global.EfficacyMap = apiEfficacyMap

	return
}
