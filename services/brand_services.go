package services

import (
	"github.com/kobeld/duoerl/models/brands"
	"github.com/kobeld/duoerl/utils"
	"github.com/kobeld/duoerlapi"
	"labix.org/v2/mgo/bson"
)

func NewBrand() (brandInput *duoerlapi.BrandInput) {
	brandInput = &duoerlapi.BrandInput{Id: bson.NewObjectId().Hex()}
	return
}

func AllBrands() (apiBrands []*duoerlapi.Brand, err error) {

	dbBrands, err := brands.FindAll(bson.M{})
	if err != nil {
		utils.PrintStackAndError(err)
		return
	}

	apiBrands = toApiBrands(dbBrands)

	return
}

func ShowBrand(brandId string) (apiBrand *duoerlapi.Brand, err error) {
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
		LogoUrl: brandInput.LogoUrl,
	}

	if err = brand.Save(); err != nil {
		utils.PrintStackAndError(err)
		return
	}

	return
}

// ------------------

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
			Name:    brand.Name,
			Alias:   brand.Alias,
			Intro:   brand.Intro,
			Country: brand.Country,
			Website: brand.Website,
			LogoUrl: brand.LogoUrl,
		}
	}

	return apiBrand
}
