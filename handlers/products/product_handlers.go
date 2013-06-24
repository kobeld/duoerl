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
	ApiProduct   *duoerlapi.Product
	ApiProducts  []*duoerlapi.Product
	Validated    *govalidations.Validated
	ApiBrands    []*duoerlapi.Brand
	ApiReviews   []*duoerlapi.Review
	ReviewInput  *duoerlapi.ReviewInput
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
	currentUserId := services.FetchAccountIdFromSession(env)

	// Get Product
	apiProduct, err := services.ShowProduct(productId, currentUserId)
	if err != nil {
		panic(err)
	}

	// Get Product Reviews
	apiReviews, err := services.ShowReviewsInProduct(productId)
	if err != nil {
		panic(err)
	}

	// Init new Review Form data
	productViewData := &ProductViewData{
		ApiProduct:  apiProduct,
		ApiReviews:  apiReviews,
		ReviewInput: services.NewReview(),
	}

	mangotemplate.ForRender(env, "products/show", productViewData)
	return
}

func New(env Env) (status Status, headers Headers, body Body) {
	productInput := services.NewProduct()
	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	productViewData := &ProductViewData{
		ProductInput: productInput,
		ApiBrands:    apiBrands,
	}

	mangotemplate.ForRender(env, "products/new", productViewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	productInput := new(duoerlapi.ProductInput)
	formdata.UnmarshalByNames(env.Request().Request, &productInput, productFields)
	productInput.AuthorId = services.FetchAccountIdFromSession(env)

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
