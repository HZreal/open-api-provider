package routers

import "open-api-interface/api"

func addNameRouter() {
	nameRouter := apiGroup.Group("user")
	{
		nameRouter.GET("name")
		nameRouter.GET("get", api.RootApi.GetName)
		nameRouter.GET("post", api.RootApi.PostName)
	}
}
