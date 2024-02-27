package api

import "github.com/gin-gonic/gin"

func (Api) GetName(c *gin.Context) {
	c.JSON(200, "get name")
}

func (Api) PostName(c *gin.Context) {
	c.JSON(200, "post name")
}
