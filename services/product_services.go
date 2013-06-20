package services

import (
	"github.com/kobeld/duoerl/models/brands"
	"github.com/kobeld/duoerl/models/products"
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

	// Collect brand Ids and find them
	brandIds := products.CollectBrandIds(dbProducts)

	dbBrands, err := brands.FindByIds(brandIds)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	// Buid the brandMap
	brandMap := brands.BuildBrandMap(dbBrands)

	apiProducts = toApiProducts(dbProducts, brandMap)

	return
}

func ShowProduct(productId string) (apiProduct *duoerlapi.Product, err error) {
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

	apiProduct = toApiProduct(product, brand)

	return
}

// Todo: validation needed
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

	product := &products.Product{
		Id:      oId,
		BrandId: brandObjectId,
		Name:    input.Name,
		Alias:   input.Alias,
		Intro:   input.Intro,
		Image:   input.Image,
	}

	if err = product.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

func toApiProducts(dbProducts []*products.Product, brandMap map[bson.ObjectId]*brands.Brand) (apiProducts []*duoerlapi.Product) {
	for _, dbProduct := range dbProducts {
		brand := brandMap[dbProduct.BrandId]
		apiProducts = append(apiProducts, toApiProduct(dbProduct, brand))
	}

	return
}

func toApiProduct(product *products.Product, brand *brands.Brand) *duoerlapi.Product {
	apiProduct := new(duoerlapi.Product)
	if product != nil {
		apiProduct = &duoerlapi.Product{
			Id:    product.Id.Hex(),
			Name:  product.Name,
			Alias: product.Alias,
			Intro: product.Intro,
			Brand: toApiBrand(brand),
		}
	}

	return apiProduct
}
