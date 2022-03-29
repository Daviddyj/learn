package main

import (
	"encoding/base64"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"learn/pkg/apis"
)

func main() {
	r := gin.Default()

	pprof.Register(r)
	r.GET("/", func(c *gin.Context) {
		c.Writer.Write([]byte(`你好，gin！`))
	})

	r.GET("/history", func(c *gin.Context) {
		c.JSON(200, []*apis.PersonInformation{
			{Name: "一杰",
				Sex:    "男",
				Tall:   1.75,
				Weight: 75,
				Age:    26},
			{Name: "瑞瑞",
				Sex:    "女",
				Tall:   1.68,
				Weight: 68,
				Age:    26},
		})
	})

	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(400, gin.H{
				"message": "无法读取个人信息"})
			return
		}
		//Todo 注册到排行榜
		c.JSON(200, "返回的内容:戴一杰是大帅逼")
	})
	r.POST("/getwithquery", func(c *gin.Context) {
		name := c.Query("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})

	r.POST("/getwithparam/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"payload": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})
	r.POST("/getwithparam/daiyijie", func(c *gin.Context) {
		name := c.Param("name")
		c.JSON(200, gin.H{
			"special": base64.StdEncoding.EncodeToString([]byte(name)),
		})
	})

	r.Run("127.0.0.1:8082")
}
