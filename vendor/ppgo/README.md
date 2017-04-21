# ppgo V 0.0.1

##前言

ppGo是一个Go语言开发的一体化开发框架,主要用于API开发,因为使用ECHO框架作为http服务,MVC模式一样可以使用,整合了比较好的组件比如echo,gorm,viper等等.

样例git:https://git.corpautohome.com/wangzeyi/ppgo-sample

**注意:框架前期还不是很完善**

##Holle,world!

创建文件 server.go

    package main
    
    import (
        "ppgo"
        "github.com/labstack/echo"
    )
    
    func main() {
    
        //初始化ECHO路由
        ppgo.NewEcho()
        // Routes路由
        ppgo.Echo.Get("/", func(c echo.Context) error {
            Response := ppgo.NewResponse(c)
            return Response.RetSuccess("hello,world!")
        })
        //开启服务
        ppgo.RunFasthttp(":1333")
    }

运行:

    go run server.go
    
请求**localhost:1333**:

![](http://i.imgur.com/tHi9dT2.png)
    
##依赖

    //配置文件读取
    go get github.com/spf13/viper
    
    //辅助使用,参数过滤,curl等(已经集成到框架)
    go get github.com/astaxie/beego
    
    //主要路由
    go get github.com/labstack/echo
    
    //主要数据操作
    go get github.com/jinzhu/gorm
    
    //log记录
    go get github.com/Sirupsen/logrus
    
    //进程级别缓存
    go get github.com/coocood/freecache
    
    //高速http
    go get github.com/valyala/fasthttp
    
    //redis依赖
    go get github.com/garyburd/redigo
    
    //注意会使用到如下依赖(国内可能需要翻墙)
    golang.org/x/net/context
    golang.org/x/sys/unix
    golang.org/x/crypto/md4

    
##联系方式

wangzeyi



