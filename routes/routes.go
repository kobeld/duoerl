package routes

import (
	"github.com/bmizerany/pat"
	hFeeds "github.com/kobeld/duoerl/handlers/feeds"
	"github.com/kobeld/duoerl/middlewares"
	"github.com/kobeld/mangogzip"
	"github.com/paulbellamy/mango"
	"github.com/sunfmin/mangolog"
	"net/http"
)

func Mux() (mux *http.ServeMux) {
	p := pat.New()
	sessionMW := mango.Sessions("f908b1c425062e95d30b8d30de7123457", "qortex",
		&mango.CookieOptions{Path: "/", MaxAge: 3600 * 24 * 7})
	rendererMW := middlewares.ProduceRenderer()
	rHtml, _ := middlewares.RespondHtml(), middlewares.RespondJson()

	mainLayoutMW := middlewares.ProduceLayout(middlewares.MAIN_LAYOUT)
	mainStack := new(mango.Stack)
	mainStack.Middleware(mangogzip.Zipper, mangolog.Logger, sessionMW, mainLayoutMW, rendererMW, rHtml)

	p.Get("/", mainStack.HandlerFunc(hFeeds.Index))

	mux = http.NewServeMux()
	mux.HandleFunc("/favicon.ico", filterUrl)
	mux.Handle("/", p)
	mux.Handle("/public/", http.FileServer(http.Dir(".")))
	return
}

func filterUrl(w http.ResponseWriter, r *http.Request) {
	return
}
