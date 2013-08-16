package services

import (
	"github.com/kobeld/duoerl/models/brands"
	"github.com/kobeld/duoerl/models/products"
	"github.com/kobeld/duoerl/models/users"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

func NewProduct() (productInput *duoerlapi.ProductInput) {
	productInput = &duoerlapi.ProductInput{Id: bson.NewObjectId().Hex()}
	return
}

func AllProducts() (apiProducts []*duoerlapi.Product, err error) {

	// Find all the products
	dbProducts, err := products.FindAll(bson.M{})
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Collect brand/author Ids and find them
	brandIds, authorIds := products.CollectBrandAndAuthorIds(dbProducts)

	dbBrands, err := brands.FindByIds(brandIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	dbAuthors, err := users.FindByIds(authorIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Build the brandMap and authorMap
	brandMap := brands.BuildBrandMap(dbBrands)
	authorMap := users.BuildUserMap(dbAuthors)

	apiProducts = toApiProducts(dbProducts, brandMap, authorMap)

	return
}

func BrandProducts(brandId string) (apiProducts []*duoerlapi.Product, err error) {

	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	productz, err := products.FindByBrandId(brandOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	for _, product := range productz {
		apiProducts = append(apiProducts, toApiProduct(product, nil, nil))
	}

	return
}

func ShowProduct(productId, userId string) (apiProduct *duoerlapi.Product, err error) {
	productOId, err := utils.ToObjectId(productId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	product, err := products.FindById(productOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand, err := brands.FindById(product.BrandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	author, err := users.FindById(product.AuthorId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiProduct = toApiProduct(product, brand, author)

	// Not login user
	if userId == "" {
		return
	}

	if wishItem, _ := GetWishItem(userId, productId); wishItem != nil {
		apiProduct.HasWished = true
	}

	if ownItem, _ := GetOwnItem(userId, productId); ownItem != nil {
		apiProduct.HasOwned = true
	}

	return
}

// Todo: validation needed. Idea: validate the input object
func CreateProduct(input *duoerlapi.ProductInput) (originInput *duoerlapi.ProductInput, err error) {
	originInput = input

	oId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandObjectId, err := utils.ToObjectId(input.BrandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	categoryOId, err := utils.ToObjectId(input.CategoryId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	subCategoryOId, err := utils.ToObjectId(input.SubCategoryId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	authorOId, err := utils.ToObjectId(input.AuthorId)
	if err != nil {
		// Don't return
		utils.PrintStackAndError(err)
	}

	product := &products.Product{
		Id:            oId,
		BrandId:       brandObjectId,
		Name:          input.Name,
		Alias:         input.Alias,
		Intro:         input.Intro,
		Image:         input.Image,
		AuthorId:      authorOId,
		CategoryId:    categoryOId,
		SubCategoryId: subCategoryOId,
		EfficacyIds:   utils.TurnPlainIdsToObjectIds(input.EfficacyIds),
	}

	if err = product.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func toApiProducts(dbProducts []*products.Product, brandMap map[bson.ObjectId]*brands.Brand,
	authorMap map[bson.ObjectId]*users.User) (apiProducts []*duoerlapi.Product) {

	for _, dbProduct := range dbProducts {
		brand := brandMap[dbProduct.BrandId]
		author := authorMap[dbProduct.AuthorId]
		apiProducts = append(apiProducts, toApiProduct(dbProduct, brand, author))
	}

	return
}

func toApiProduct(product *products.Product, brand *brands.Brand, author *users.User) *duoerlapi.Product {
	apiProduct := new(duoerlapi.Product)
	if product != nil {
		apiProduct = &duoerlapi.Product{
			Id:          product.Id.Hex(),
			Link:        product.Link(),
			Name:        product.Name,
			Alias:       product.Alias,
			Intro:       product.Intro,
			Brand:       toApiBrand(brand),
			Author:      toApiUser(author),
			Category:    GetCategory(product.CategoryId.Hex()),
			SubCategory: GetSubCategory(product.SubCategoryId.Hex()),
			Efficacies:  GetEfficaciesByIds(utils.TurnObjectIdToPlainIds(product.EfficacyIds)),
		}
	}

	return apiProduct
}
