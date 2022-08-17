# 简单的 gin 多语言支持中间件

[English](README.md) | 简体中文

发现网上的 `gin` 多语言中间件都太复杂了，于是自己动手写了个简单的。

## 示例

```go
// example/main.go
package main

import (
	"log"
	"net/http"
	"path"

	ginI18n "github.com/fishjar/gin-i18n"
	"github.com/gin-gonic/gin"
)

func main() {
	const defaultLang = "zh-CN"                     // 默认语言
	const supportLang = "zh-CN,en-US"               // 支持的语言列表 ​​(必须包含默认语言)
	var filePath = path.Join("example", "localize") // 多语言文件所在目录

	// 初始化设置
	ginI18n.LocalizerInit(defaultLang, supportLang, filePath)

	// new gin engine
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// 应用中间件
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

```yaml
# example/localize/zh-CN.yml
welcome: 欢迎！
hello_world: 你好 %s!
```

```yaml
# example/localize/en-US.yml
welcome: welcome!
hello_world: hello %s!
```

```sh
# 检查效果
go run example/main.go

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN'             # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: en-US'             # welcome!

curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN,en-US;q=0.9' # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: zh-CN;q=0.9,en-US' # welcome!

curl http://127.0.0.1:8080/ -H 'accept-language: zh'                # 欢迎！
curl http://127.0.0.1:8080/ -H 'accept-language: en'                # welcome!

curl http://127.0.0.1:8080/gabe -H 'accept-language: zh-CN'         # 你好 gabe!
curl http://127.0.0.1:8080/gabe -H 'accept-language: en-US'         # hello gabe!
```
