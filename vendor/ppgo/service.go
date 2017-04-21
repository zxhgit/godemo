package ppgo
import (
"github.com/labstack/echo"
)

func PpgoRun() {

	//初始化ECHO路由
	NewEcho()
	// Routes路由
	Echo.Get("/", func(c echo.Context) error {
		Response := NewResponse(c)
		return Response.RetSuccess("hello,world!")
	})
	//开启服务
	RunFasthttp(":1333")
}