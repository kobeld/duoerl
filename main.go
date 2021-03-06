package main

import (
	"github.com/kobeld/duoerl/configs"
	"github.com/kobeld/duoerl/global"
	"github.com/kobeld/duoerl/routes"
	"github.com/shaoshing/train"
	"github.com/sunfmin/mangotemplate"
	"github.com/sunfmin/mgodb"
	"log"
	"net/http"
)

func main() {
	mangotemplate.AutoReload = true
	train.Config.SASS.DebugInfo = false

	mgodb.Setup(configs.DBUrl, configs.DatabaseName)
	global.ImageDatabase = mgodb.NewDatabase(configs.DBUrl, configs.ImageDatabaseName)

	mux := routes.Mux()

	log.Printf("Starting server on %s\n", configs.HttpPort)
	panic(http.ListenAndServe(configs.HttpPort, mux))
}
