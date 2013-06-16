package products

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
)

func Index(env Env) (status Status, headers Headers, body Body) {

	mangotemplate.ForRender(env, "products/index", nil)
	return
}
