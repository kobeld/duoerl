package main

import (
	"github.com/kobeld/duoerl/configs"
	"github.com/kobeld/duoerl/routes"
	"github.com/sunfmin/mgodb"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	mgodb.Setup(configs.DBUrl, configs.Database)
	rand.Seed(time.Now().UnixNano())
	configs.AssetsVersion = rand.Intn(100000)

	mux := routes.Mux()

	log.Printf("Starting server on %s\n", configs.HttpPort)
	err := http.ListenAndServe(configs.HttpPort, mux)
	if err != nil {
		panic("Http ListenAndServe: " + err.Error())
	}
}
