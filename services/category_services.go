package services

import (
	"github.com/kobeld/duoerl/models/categories"
	"github.com/kobeld/duoerlapi"
)

func GetCategory(categoryId string) *duoerlapi.Category {
	return GetCategoryMap()[categoryId]
}

func GetSubCategory(subCategoryId string) *duoerlapi.SubCategory {
	return GetSubCategoryMap()[subCategoryId]
}

// ---- Private -----

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
