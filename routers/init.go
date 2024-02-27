package routers

import "github.com/gin-gonic/gin"

var Router *gin.Engine
var apiGroup *gin.RouterGroup

func init() {
	r := gin.Default()
	apiGroup = r.Group("api")
	addNameRouter()
	addSentenceRouter()
	Router = r
}
