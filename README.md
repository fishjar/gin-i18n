# Simple gin i18n middleware

English | [简体中文](README.zh-CN.md)

I found that other people's `gin` i18n middleware too complicated, so I wrote a simple one.

## Example

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
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	// apply i18n middleware
	router.Use(ginI18n.Localizer(&ginI18n.Options{
		DefaultLang:  "zh-CN",                          // default language
		SupportLangs: "zh-CN,en-US",                    // list of supported languages ​​(must include default language)
		FilePath:     path.Join("example", "localize"), // multilingual file directory
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, ginI18n.Msg(c, "welcome"))
	})

	router.GET("/:name", func(c *gin.Context) {
		c.String(http.StatusOK, ginI18n.Msg(c, "hello_world", c.Param("name")))
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
# review
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
