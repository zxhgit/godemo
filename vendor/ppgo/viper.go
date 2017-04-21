//	ppgo-Config
//	使用spf13大神的viper配置文件获取工具作为phalgo的配置文件工具

//  依赖情况:
//          "github.com/spf13/viper"

package ppgo

import (
	"github.com/spf13/viper"
	"path"
)

var Config *viper.Viper

//初始化配置文件
func NewConfig(filePath string, fileName string) {

	Config = viper.New()
	Config.WatchConfig()
	Config.SetConfigName(fileName)
	//filePath支持相对路径和绝对路径 etc:"/a/b" "b" "./b"
	if (filePath[:1] != "/"){
		Config.AddConfigPath(path.Join(GetPath(),filePath))
		//println(path.Join(GetPath(),filePath));
	}else{
		Config.AddConfigPath(filePath)
		//println(filePath);
	}

	// 找到并读取配置文件并且 处理错误读取配置文件
	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}

}


