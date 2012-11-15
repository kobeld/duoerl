package feeds

import (
	. "github.com/paulbellamy/mango"
	"github.com/sunfmin/mangotemplate"
)

func Index(env Env) (status Status, headers Headers, body Body) {

	mangotemplate.ForRender(env, "feeds/index", nil)

	return
}
