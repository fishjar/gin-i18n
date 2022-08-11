# 简单的 gin 多语言支持中间件

发现都网上的 `gin` 多语言中间件太复杂了，于是自己动手写了个简单的。

## Example

```go
package main

import (
	"log"
	"net/http"
	"path"

	ginI18n "github.com/fishjar/gin-i18n"
	"github.com/gin-gonic/gin"
)

func main() {
	const defaultLang = "zh-CN"                           // 默认语言
	const supportLang = "zh-CN,en-US"                     // 支持的语言列表(必须包含默认语言)
	var filePath = path.Join("./", "example", "localize") // 语言文件目录

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
```

```sh
# 测试效果
go run example/main.go

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN'             # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: en-US'             # welcome!

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN,en-US;q=0.9' # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN;q=0.9,en-US' # welcome!
curl http://127.0.0.1:8080/ -H 'accept-language: en'                # 欢迎！

curl http://127.0.0.1:8080/gabe -H 'accept-language: zh-CN'         # 你好 gabe!
curl http://127.0.0.1:8080/gabe -H 'accept-language: en-US'         # hello gabe!
```
