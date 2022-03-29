package main

import (
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"learn/chapter12/02.pratice/frinterface"
	"learn/pkg/apis"
	"log"
	"net/http"
)

//func main() {
//	r := gin.Default()
//
//	pprof.Register(r)
//	var rankServer frinterface.ServerInterface = NewFatRateRank()
//	r.POST("/register", func(c *gin.Context) {
//		pi := &apis.PersonInformation{}
//		if err := c.BindJSON(pi); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "无法解析注册信息"})
//			return
//		}
//		//Todo 注册到排行榜
//		if err := rankServer.RegisterPersonInformation(pi); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "注册失败"})
//			return
//		}
//		c.JSON(200, gin.H{
//			"message": "success",
//		})
//	})
//
//	r.PUT("/personalinfo", func(c *gin.Context) {
//		pi := &apis.PersonInformation{}
//		if err := c.BindJSON(pi); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "无法解析注册信息"})
//			return
//		}
//		if resp, err := rankServer.UpdatePersonInformation(pi); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "服务器更新用户失败"})
//			return
//		} else {
//			c.JSON(200, resp)
//		}
//
//	})
//
//	r.GET("/rank/:name", func(c *gin.Context) {
//		name := c.Param("name")
//		if rank, err := rankServer.GetFateRate(name); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "获取用户排名失败"})
//			return
//		} else {
//			c.JSON(200, rank)
//		}
//	})
//
//	r.GET("/ranktop", func(c *gin.Context) {
//		if frTop, err := rankServer.GetTop(); err != nil {
//			c.JSON(http.StatusBadRequest, gin.H{
//				"errMessage": "获取用户前十排名失败"})
//			return
//		} else {
//			c.JSON(200, frTop)
//		}
//	})
//
//	if err := r.Run("127.0.0.1:8081"); err != nil {
//		log.Fatal(err)
//	}
//}

func connectDb() *gorm.DB {
	conn, err := gorm.Open(mysql.Open("root:1165207594dyj@tcp(127.0.0.1:3306)/testdb"))
	if err != nil {
		log.Fatal("数据库连接失败", err)
	}
	fmt.Println("连接数据库成功")
	return conn
}

func main() {
	conn := connectDb()
	var rankServer frinterface.ServerInterface = NewDbRank(conn, NewFatRateRank())
	if initFank, ok := rankServer.(frinterface.RankInitInterface); ok {
		if err := initFank.Init(); err != nil {
			log.Fatal("初始化失败:", err)
		}
	}
	r := gin.Default()

	pprof.Register(r)

	r.POST("/register", func(c *gin.Context) {
		pi := &apis.PersonInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "无法解析注册信息"})
			return
		}
		//Todo 注册到排行榜
		if err := rankServer.RegisterPersonInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "注册失败"})
			return
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.PUT("/personalinfo", func(c *gin.Context) {
		pi := &apis.PersonInformation{}
		if err := c.BindJSON(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "无法解析注册信息"})
			return
		}
		if resp, err := rankServer.UpdatePersonInformation(pi); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "服务器更新用户失败"})
			return
		} else {
			c.JSON(200, resp)
		}

	})

	r.GET("/rank/:name", func(c *gin.Context) {
		name := c.Param("name")
		if rank, err := rankServer.GetFateRate(name); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "获取用户排名失败"})
			return
		} else {
			c.JSON(200, rank)
		}
	})

	r.GET("/ranktop", func(c *gin.Context) {
		if frTop, err := rankServer.GetTop(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"errMessage": "获取用户前十排名失败"})
			return
		} else {
			c.JSON(200, frTop)
		}
	})

	if err := r.Run("127.0.0.1:8081"); err != nil {
		log.Fatal(err)
	}
}
