package products

import (
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/govalidations"
	"github.com/sunfmin/mangotemplate"
	"github.com/theplant/formdata"
	"net/http"
)

var (
	productFields = []string{"Id", "Name", "Alias", "Intro", "Image", "BrandId",
		"CategoryId", "SubcategoryId", "EfficacyIds"}
)

type ProductViewData struct {
	ProductInput   *duoerlapi.ProductInput
	ApiProduct     *duoerlapi.Product
	ApiProducts    []*duoerlapi.Product
	ApiCategories  []*duoerlapi.Category
	Validated      *govalidations.Validated
	ApiBrands      []*duoerlapi.Brand
	ApiReviews     []*duoerlapi.Review
	ReviewInput    *duoerlapi.ReviewInput
	GotFromOptions map[string]string
	RatingOptions  map[string]string
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
	currentUserId := services.FetchUserIdFromSession(env)

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
		ApiProduct:     apiProduct,
		ApiReviews:     apiReviews,
		ReviewInput:    services.NewReview(),
		GotFromOptions: global.GotFromOptions,
		RatingOptions:  global.RatingOptions,
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
		ProductInput:  productInput,
		ApiBrands:     apiBrands,
		ApiCategories: services.GetCategories(),
	}

	mangotemplate.ForRender(env, "products/new", productViewData)
	return
}

func Edit(env Env) (status Status, headers Headers, body Body) {

	productId := env.Request().URL.Query().Get(":id")

	productInput, err := services.EditProduct(productId)
	if err != nil {
		panic(err)
	}

	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	productViewData := &ProductViewData{
		ProductInput:  productInput,
		ApiBrands:     apiBrands,
		ApiCategories: services.GetCategories(),
	}

	mangotemplate.ForRender(env, "products/edit", productViewData)
	return
}

func Create(env Env) (status Status, headers Headers, body Body) {

	productInput := new(duoerlapi.ProductInput)
	formdata.UnmarshalByNames(env.Request().Request, &productInput, productFields)
	productInput.AuthorId = services.FetchUserIdFromSession(env)

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

func Update(env Env) (status Status, headers Headers, body Body) {

	productInput := new(duoerlapi.ProductInput)
	formdata.UnmarshalByNames(env.Request().Request, &productInput, productFields)

	result, err := services.UpdateProduct(productInput)
	if validated, ok := err.(*govalidations.Validated); ok {
		viewData := &ProductViewData{
			ProductInput: result,
			Validated:    validated,
		}
		mangotemplate.ForRender(env, "products/edit", viewData)
		return
	}
	if err != nil {
		panic(err)
	}

	return Redirect(http.StatusFound, "/product/"+result.Id)
}
