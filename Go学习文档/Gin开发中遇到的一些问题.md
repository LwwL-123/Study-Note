# Gin开发中遇到的一些问题



1. get和post的form表单中，通过url提交时，gin会自动把特殊符号进行转义

```
curl -g "http://localhost:8080/pk/queryKey" -H "Content-Type: application/json" -d {\"public_key\":\""ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDHmVkaKCU9lkeYB4bScoyrjGFk0M9wYFwfFEI1lXsWYCjuFWGtIwd8wHT+ApGsUY83y/5AWjxRhO6yNjK+9mDE+UxLYNvaDhWEsw9UlXCKLg7gcDgZmdLkr5ellJjuW5crsMWMwCg7CNigL6vxCTLu5lzgbeqJT09W7/hhyn7pcFw7ggul5O530hhvv+rvhZtkurHVOUHp8ofiSZdOtmp8ZgDtsP6CzsAUgiyBSswaP4TFyrn+USbr9Vt9A+qhCBa2RT2ADPY3jEPvslf9wnzYsiROdzneb9AB+HCEsr53CeGHAym1DcW0wZ2Nwf2k0rXVtyCUt1O2zptkbhl5Gav5PYYhsC0TYUJdUDT1ZELv4Uqbs7Nw8ILOk7PY/R6wcu3ZxBqOdbGc7FKRbaPkPzSC3ozt8yQ2S7WLDTTNgHICKXBYTCDFFCEcV9d0MxMwvIGZ67MMKb71Udmh1BdlBODbNQdEtXea07GkFHZ6bjbR2oOOBEuoXwWzYx9JtJiolrM= lzw@MacBook-Air.local"\"}
```



解决：采用json的方式提交，通过post的body请求体



2. 需要传入一个新的参数给gin的handler

解决：

```
// 定义一个新的context
type NewContext struct {
	c   *gin.Context
	ctx context.Context
}
type HandlerFunc func(*NewContext)

// 定义handler
func handler(handler HandlerFunc, ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		newContext := new(NewContext)
		newContext.c = c
		newContext.ctx = ctx
		handler(newContext)
	}
}

	// 调用
	pk.POST("/grantKey", handler(grantKey, ctx))

```