# 简单的 gin 多语言支持中间件

发现都网上的 `gin` 多语言中间件太复杂了，于是自己动手写了个简单的。

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
