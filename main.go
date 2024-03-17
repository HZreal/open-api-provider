package main

import (
	"open-api-provider/config"
	"open-api-provider/routers"
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
