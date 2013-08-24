package services

import (
	"github.com/kobeld/duoerl/models/brands"
	"github.com/kobeld/duoerl/models/followbrands"
	"github.com/kobeld/duoerl/models/images"
	"github.com/kobeld/duoerl/models/products"
	"github.com/kobeld/duoerl/models/reviews"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

func NewBrand() (brandInput *duoerlapi.BrandInput) {
	brandInput = &duoerlapi.BrandInput{
		Id:        bson.NewObjectId().Hex(),
		Logo:      "http://lorempixel.com/g/200/200/", // Temp
		ImageAttr: newBrandImageAttr(),
	}
	return
}

func EditBrand(brandId string) (brandInput *duoerlapi.BrandInput, err error) {

	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand, err := brands.FindById(brandOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brandInput = toBrandInput(brand)

	return
}

func UpdateBrand(input *duoerlapi.BrandInput) (originInput *duoerlapi.BrandInput, err error) {
	originInput = input

	brandOId, err := utils.ToObjectId(input.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand, err := brands.FindById(brandOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand.Name = input.Name
	brand.Alias = input.Alias
	brand.Intro = input.Intro
	brand.Country = input.Country
	brand.Logo = input.Logo
	brand.Website = input.Website

	if err = brand.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

// Need to be cached
func AllBrands() (apiBrands []*duoerlapi.Brand, err error) {

	dbBrands, err := brands.FindAll(bson.M{})
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiBrands = toApiBrands(dbBrands)

	return
}

func ShowBrand(brandId, userId string) (apiBrand *duoerlapi.Brand, err error) {
	brandOId, err := utils.ToObjectId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand, err := brands.FindById(brandOId)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiBrand = toApiBrand(brand)

	// Not login user
	if userId == "" {
		return
	}

	if followBrand := GetFollowBrand(userId, brandId); followBrand != nil {
		apiBrand.HasFollowed = true
	}

	apiBrand.BrandStats = getBrandStats(brandOId)

	return
}

func CreateBrand(brandInput *duoerlapi.BrandInput) (input *duoerlapi.BrandInput, err error) {
	input = brandInput

	oId, err := utils.ToObjectId(brandInput.Id)
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	brand := &brands.Brand{
		Id:      oId,
		Name:    brandInput.Name,
		Alias:   brandInput.Alias,
		Intro:   brandInput.Intro,
		Country: brandInput.Country,
		Website: brandInput.Website,
		Logo:    brandInput.Logo,
	}

	if err = brand.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

// -------- Private ----------

func getBrandStats(brandId bson.ObjectId) (brandStats *duoerlapi.BrandStats) {
	var err error
	brandStats = new(duoerlapi.BrandStats)

	brandStats.FollowerCount, err = followbrands.CountBrandFollowerByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
	}

	brandStats.ProductCount, err = products.CountProductByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
	}

	brandStats.ReviewCount, err = reviews.CountReviewByBrandId(brandId)
	if err != nil {
		utils.PrintStackAndError(err)
	}

	return brandStats
}

func newBrandImageAttr() *duoerlapi.ImageAttr {
	return &duoerlapi.ImageAttr{
		ImageType: images.CATEGORY_BRAND,
	}
}

func toApiBrands(dbBrands []*brands.Brand) (apiBrands []*duoerlapi.Brand) {
	for _, dbBrand := range dbBrands {
		apiBrands = append(apiBrands, toApiBrand(dbBrand))
	}
	return
}

func toApiBrand(brand *brands.Brand) *duoerlapi.Brand {
	apiBrand := new(duoerlapi.Brand)

	if brand != nil {
		apiBrand = &duoerlapi.Brand{
			Id:      brand.Id.Hex(),
			Link:    brand.Link(),
			Name:    brand.Name,
			Alias:   brand.Alias,
			Intro:   brand.Intro,
			Country: brand.Country,
			Website: brand.Website,
			Logo:    brand.LogoUrl(),
		}
	}

	return apiBrand
}

func toBrandInput(brand *brands.Brand) (brandInput *duoerlapi.BrandInput) {
	brandInput = &duoerlapi.BrandInput{
		Id:        brand.Id.Hex(),
		Name:      brand.Name,
		Alias:     brand.Alias,
		Intro:     brand.Intro,
		Country:   brand.Country,
		Website:   brand.Website,
		Logo:      brand.LogoUrl(),
		ImageAttr: newBrandImageAttr(),
	}

	return
}
