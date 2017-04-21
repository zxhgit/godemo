package main
import (
	"fmt"
	//"math/rand"
	//"math"
	//"math/cmplx"
	//"math"
	//"runtime"
	//"time"
	//"io"
	//"strings"
	//"os"
	"strconv"
	"flag"
	"os"
	//"lib1"
	//"lib2"
	"ppgo"
	"github.com/labstack/echo"
)
func main() {
	//primeTest()
	//func1.DoTest()
	//func2.DoTest2()

	//ppgo.PpgoRun()


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

var goal int
func primetask(c chan int) {
	p := <-c
	if p > goal {
		os.Exit(0)
	}
	fmt.Println(p)
	nc := make(chan int)
	go primetask(nc)
	for {
		i := <-c
		if i%p != 0 {
			nc <- i
		}
	}
}
func primeTest() {
	flag.Parse()
	args := flag.Args()
	if args != nil&&len(args) > 0 {
		var err error
		goal, err = strconv.Atoi(args[0])
		if err != nil {
			goal = 100
		}
	} else {
		goal = 100
	}
	fmt.Println("goal=", goal)
	c := make(chan int)
	go primetask(c)
	for i := 2;; i++ {
		c <- i
	}
}
