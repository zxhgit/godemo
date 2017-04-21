package test

import (
	"fmt"
	"testing"
	"ppgo"
)


func TestConfig(t *testing.T) {

	ppgo.API_ROOT = "/Users/wangpp/Code/github/go/src/ppgo-sample";
	//初始化配置文件
	ppgo.NewConfig("Config", "conf")

	fmt.Println(ppgo.Config.GetString("system.port"));
}


