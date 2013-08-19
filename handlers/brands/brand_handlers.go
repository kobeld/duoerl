package brands

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"net/http"
)

var (
	brandFields = []string{"Id", "Name", "Alias", "Country",
		"Intro", "Website", "Logo"}
)

type BrandViewData struct {
	BrandInput  *duoerlapi.BrandInput
	Validated   *govalidations.Validated
	ApiBrand    *duoerlapi.Brand
	ApiBrands   []*duoerlapi.Brand
	ApiProducts []*duoerlapi.Product
}

func newBrandViewData(brandInput *duoerlapi.BrandInput,
	validated *govalidations.Validated) *BrandViewData {

	return &BrandViewData{
		BrandInput: brandInput,
		Validated:  validated,
	}
}

// ----------------

func Index(env Env) (status Status, headers Headers, body Body) {

	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "brands/index", &BrandViewData{ApiBrands: apiBrands})
	return
}

func Show(env Env) (status Status, headers Headers, body Body) {
	brandId := env.Request().URL.Query().Get(":id")
	currentUserId := services.FetchUserIdFromSession(env)

	apiBrand, err := services.ShowBrand(brandId, currentUserId)
	if err != nil {
		panic(err)
	}

	apiProducts, err := services.BrandProducts(brandId)

	brandViewData := &BrandViewData{
		ApiBrand:    apiBrand,
		ApiProducts: apiProducts,
	}

	mangotemplate.ForRender(env, "brands/show", brandViewData)
	return
}

func New(env Env) (status Status, headers Headers, body Body) {
	brandInput := services.NewBrand()
	mangotemplate.ForRender(env, "brands/new", newBrandViewData(brandInput, nil))
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	brandInput := new(duoerlapi.BrandInput)
	formdata.UnmarshalByNames(env.Request().Request, &brandInput, brandFields)

	result, err := services.CreateBrand(brandInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		mangotemplate.ForRender(env, "brands/new", newBrandViewData(result, validated))
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/brand/"+result.Id)
}

func Edit(env Env) (status Status, headers Headers, body Body) {

	brandId := env.Request().URL.Query().Get(":id")
	brandInput, err := services.EditBrand(brandId)
	if validated, ok := err.(*govalidations.Validated); ok {
		mangotemplate.ForRender(env, "brands/edit", newBrandViewData(brandInput, validated))
		return
	}
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "brands/edit", newBrandViewData(brandInput, nil))
	return
}

func Update(env Env) (status Status, headers Headers, body Body) {

	brandInput := new(duoerlapi.BrandInput)
	formdata.UnmarshalByNames(env.Request().Request, &brandInput, brandFields)

	result, err := services.UpdateBrand(brandInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		mangotemplate.ForRender(env, "products/edit", newBrandViewData(result, validated))
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/brand/"+result.Id)
}
