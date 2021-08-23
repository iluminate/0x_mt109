package main

import (
	"0x_mt109/application/configs"
	"0x_mt109/application/routers"
	"0x_mt109/cmd/api/server"
	"0x_mt109/helpers/loader"
)

func main() {
	serverConfig := new(configs.ConfigServer)
	loader.ReadConf(serverConfig)
	server := server.NewHttpServer(serverConfig)
	router := routers.NewHttpRouter(serverConfig.Server.ContextPath)
	router.EnableCORS("*")
	server.ListenAndServe(router.Handler())
}