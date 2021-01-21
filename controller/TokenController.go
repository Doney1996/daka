package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

type TokenContent struct {
	sync.Mutex
	Token string `json:"token"`
}

var token = new(TokenContent)

func GetToken(c *gin.Context) {
	//jsonBytes, _ := json.Marshal(token) //转换成JSON返回的是byte[]
	c.JSON(http.StatusOK, token)
}
func PostToken(c *gin.Context) {
	jsonContent := make(map[string]string) //注意该结构接受的内容
	_ = c.BindJSON(&jsonContent)
	tt := jsonContent["token"]
	if len(tt) < 1 {
		c.JSON(http.StatusOK, "token不能为空")
		return
	}
	submitToken(tt)
	c.JSON(http.StatusOK, "token提交成功")
}
func submitToken(tt string) {
	token.Lock()
	token.Token = tt
	token.Unlock()
}
