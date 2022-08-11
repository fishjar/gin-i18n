package main

import (
	"log"
	"net/http"
	"path"

	ginI18n "github.com/fishjar/gin-i18n"
	"github.com/gin-gonic/gin"
)

func main() {
	const defaultLang = "zh-CN"                          // 默认语言
	const supportLang = "zh-CN,en-US"                    // 支持的语言列表(必须包含默认语言)
	var filePath = path.Join(".", "example", "localize") // 语言文件目录

	// localizer init
	ginI18n.LocalizerInit(defaultLang, supportLang, filePath)

	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(ginI18n.GinLocalizer())

	router.GET("/", func(c *gin.Context) {
		localizer := c.MustGet("Localizer").(*ginI18n.UserLocalize)
		c.String(http.StatusOK, localizer.GetMsg("welcome"))
	})

	router.GET("/:name", func(c *gin.Context) {
		localizer := c.MustGet("Localizer").(*ginI18n.UserLocalize)
		c.String(http.StatusOK, localizer.GetMsg("hello_world", c.Param("name")))
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
