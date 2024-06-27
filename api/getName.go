package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"open-api-provider/utils"
)

func (Api) GetNameByGet(c *gin.Context) {
	name := c.Query("name")
	c.JSON(200, gin.H{"name": name})
}

type GetNameReqHeader struct {
	AppKey     string `json:"appKey"`
	Parameters string `json:"parameters"`
	Timestamp  string `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Sign       string `json:"sign"`
}

type GetNameReqData struct {
	Name string `json:"name"`
}

func (Api) GetNameByPost(c *gin.Context) {
	// 解析 Header
	var reqHeader GetNameReqHeader
	if err := c.ShouldBindHeader(&reqHeader); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty header"})
		return
	}

	// TODO 判断重放
	// redis 查询是否存在 nonce，存在则拒绝，否则写入

	// TODO 从数据库查询 appKey 和 appSecret, 先模拟一下
	appSecret := "asdfgqwertzxcvb"
	if reqHeader.AppKey != "qwert123456" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No appKey"})
	}

	// 验签
	serverSign := utils.GenSign(reqHeader.Parameters, appSecret)
	if serverSign != reqHeader.Sign {
		c.JSON(http.StatusBadRequest, gin.H{"error": "sign error"})
	}

	// 解析 JSON 请求体
	var requestData GetNameReqData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 返回 JSON 响应
	c.JSON(http.StatusOK, gin.H{"name": requestData.Name})
}
