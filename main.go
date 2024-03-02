package main

import (
	"open-api-interface/config"
	"open-api-interface/routers"
)

func main() {
	//
	ginConfig := config.Conf.Gin.GetGinAddr()

	//
	err := routers.Router.Run(ginConfig)
	if err != nil {
		return
	}
}
