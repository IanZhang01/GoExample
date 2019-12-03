package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getinfos(c *gin.Context) {
	// post man
	// url localhost:xxxx?name=xxx
	name := c.Query("name")
	lastname := c.DefaultQuery("lastname", "這是預設值")
	c.String(http.StatusOK, "Hello %s,%s", name, lastname)
}

func postinfos(c *gin.Context) {
	// post man
	// x-www-form-urlencoded data format
	name := c.PostForm("name")
	lastname := c.DefaultPostForm("lastname", "這是預設值")
	c.String(http.StatusOK, "Hello %s,%s", name, lastname)
}

// Login blablabla
type Login struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

// json output
func loginJSON(c *gin.Context) {
	json := Login{}
	//獲取json資料後並解析
	if c.BindJSON(&json) == nil {
		if json.Name == "root" && json.Password == "root" {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "賬號或者密碼錯誤"})
		}
	}
}

// FormLogin bababa
type FormLogin struct {
	Name     string `form:"name"`
	Password string `form:"password"`
}

func loginFORM(c *gin.Context) {
	form := Login{}
	//獲取form資料後並解析
	if c.Bind(&form) == nil {
		if form.Name == "root" && form.Password == "root" {
			c.JSON(200, gin.H{"status": "登陸成功"})
		} else {
			c.JSON(203, gin.H{"status": "賬號或者密碼錯誤"})
		}
	}
}

func main() {
	r := gin.Default()
	r.GET("/ping", getinfos)
	r.POST("/postinfos", postinfos)
	r.POST("/loginJSON", loginJSON)
	r.POST("/loginFORM", loginFORM)
	r.Run(":7211")
}
