package brands

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/formdata"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"net/http"
)

var (
	entryFields = []string{"Id", "Name", "Alias", "Country", "Intro", "Website", "LogoUrl"}
)

type BrandViewData struct {
	BrandInput *duoerlapi.BrandInput
	Validated  *govalidations.Validated
	ApiBrands  []*duoerlapi.Brand
}

func NewBrandViewData(brandInput *duoerlapi.BrandInput, validated *govalidations.Validated) *BrandViewData {
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

	apiBrand, err := services.ShowBrand(brandId)
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "brands/show", apiBrand)
	return
}

func New(env Env) (status Status, headers Headers, body Body) {
	brandInput := services.NewBrand()
	mangotemplate.ForRender(env, "brands/new", brandInput)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	brandInput := new(duoerlapi.BrandInput)
	formdata.UnmarshalByNames(env.Request().Request, &brandInput, entryFields)

	result, err := services.CreateBrand(brandInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		mangotemplate.ForRender(env, "brands/new", NewBrandViewData(result, validated))
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/brand/"+result.Id)
}

func Edit(env Env) (status Status, headers Headers, body Body) {

	// mangotemplate.ForRender(env, template, data)
	return
}

func Update(env Env) (status Status, headers Headers, body Body) {

	// mangotemplate.ForRender(env, template, data)
	return
}
