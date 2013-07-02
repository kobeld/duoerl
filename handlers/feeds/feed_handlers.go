package feeds

import (
	"github.com/kobeld/duoerl/services"
	"github.com/kobeld/duoerlapi"
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
)

type HomeViewData struct {
	ApiProducts []*duoerlapi.Product
	ApiBrands   []*duoerlapi.Brand
}

func Index(env Env) (status Status, headers Headers, body Body) {

	apiProducts, err := services.AllProducts()
	if err != nil {
		panic(err)
	}

	apiBrands, err := services.AllBrands()
	if err != nil {
		panic(err)
	}

	homeViewData := &HomeViewData{
		ApiProducts: apiProducts,
		ApiBrands:   apiBrands,
	}

	mangotemplate.ForRender(env, "feeds/index", homeViewData)

	return
}
