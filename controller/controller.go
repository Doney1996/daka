package controller

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

/*
 url     --> controller  --> logic   -->    model
请求来了  -->  控制器      --> 业务逻辑  --> 模型层的增删改查
*/

func IndexHandler(c *gin.Context) {
	c.Redirect(200, "/index.html")
	//c.HTML(http.StatusOK, "index.html", nil)
}

func Daka(c *gin.Context) {
	jsonContent := make(map[string]string) //注意该结构接受的内容
	_ = c.BindJSON(&jsonContent)
	token := jsonContent["token"]
	uid := jsonContent["uid"]
	longi := jsonContent["longi"]
	lati := jsonContent["lati"]
	address := jsonContent["address"]
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, "token不能为空")
		return
	}
	if len(uid) == 0 {
		c.JSON(http.StatusBadRequest, "uid不能为空")
		return
	}
	if len(longi) == 0 {
		c.JSON(http.StatusBadRequest, "longi经度不能为空")
		return
	}
	if len(lati) == 0 {
		c.JSON(http.StatusBadRequest, "lati维度不能为空")
		return
	}
	if len(address) == 0 {
		c.JSON(http.StatusBadRequest, "address地址不能为空")
		return
	}
	postJson := []byte(fmt.Sprintf(`{
    "access_token":"%s",
    "base_param":"ehr01",
    "accountid":"%s",
    "point":"%s,%s,%s",
    "photo":"",
    "bz":"",
    "wifi":"",
    "qqlx":"1",
    "protocalMustParams":{
        "baseParam":{
            "userCode":"nbzlb"
        },
        "charset":"utf-8",
        "appid":"ehradmin"
    }
    }`, token, uid, longi, lati, address))
	var jsonStr = []byte(postJson)
	buf := bytes.NewBuffer(jsonStr)
	get, _ := http.Post("https://epp.epsoft.com.cn/ehr/signServlet", "application/json;charset=UTF-8", buf)
	defer get.Body.Close()
	body, _ := ioutil.ReadAll(get.Body)
	c.JSON(http.StatusOK, string(body))

}

//查询打卡信息
func Chaxun(c *gin.Context) {
	jsonContent := make(map[string]string) //注意该结构接受的内容
	c.BindJSON(&jsonContent)
	token := jsonContent["token"]
	uid := jsonContent["uid"]
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, "token不能为空")
		return
	}
	if len(uid) == 0 {
		c.JSON(http.StatusBadRequest, "uid不能为空")
		return
	}
	postJson := []byte(fmt.Sprintf(`
        {
    "access_token":"%s",
    "base_param":"comm07",
    "accountid":"%s",
    "qqlx":"1",
    "protocalMustParams":{
        "baseParam":{
            "userCode":"nbzlb"
        },
        "charset":"utf-8",
        "appid":"ehradmin"
    }
}`, token, uid))
	var jsonStr = []byte(postJson)
	buf := bytes.NewBuffer(jsonStr)
	get, _ := http.Post("https://epp.epsoft.com.cn/ehr/CommonInfoServlet", "application/json;charset=UTF-8", buf)
	defer get.Body.Close()
	body, _ := ioutil.ReadAll(get.Body)
	c.JSON(http.StatusOK, string(body))

}
