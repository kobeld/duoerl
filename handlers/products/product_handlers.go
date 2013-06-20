package products

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
	productFields = []string{"Id", "Name", "Alias", "Intro", "Image", "BrandId"}
)

type ProductViewData struct {
	ProductInput *duoerlapi.ProductInput
	ApiProducts  []*duoerlapi.Product
	Validated    *govalidations.Validated
	Brands       []*duoerlapi.Brand
}

func newProductViewData(productInput *duoerlapi.ProductInput,
	brands []*duoerlapi.Brand) *ProductViewData {

	return &ProductViewData{
		ProductInput: productInput,
		Brands:       brands,
	}
}

// ----------------

func Index(env Env) (status Status, headers Headers, body Body) {

	apiProducts, err := services.AllProducts()
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "products/index", &ProductViewData{ApiProducts: apiProducts})
	return
}

func Show(env Env) (status Status, headers Headers, body Body) {
	productId := env.Request().URL.Query().Get(":id")

	apiProduct, err := services.ShowProduct(productId)
	if err != nil {
		panic(err)
	}

	mangotemplate.ForRender(env, "products/show", apiProduct)
	return
}

func New(env Env) (status Status, headers Headers, body Body) {
	productInput := services.NewProduct()
	brands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	productViewData := newProductViewData(productInput, brands)
	mangotemplate.ForRender(env, "products/new", productViewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	productInput := new(duoerlapi.ProductInput)
	formdata.UnmarshalByNames(env.Request().Request, &productInput, productFields)

	result, err := services.CreateProduct(productInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData := &ProductViewData{
			ProductInput: result,
			Validated:    validated,
		}
		mangotemplate.ForRender(env, "products/new", viewData)
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/product/"+result.Id)
}
