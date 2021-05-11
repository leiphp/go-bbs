/*
	提供不同环境配置文件，提供配置文件读写
*/
package configs

import (
	"flag"
	"github.com/spf13/viper"
)

var EnvPort *string

//	返回 *viper.Viper对象
func InitConfig() *viper.Viper {
	//取得viper对象
	Yaml := viper.New()
	//启动go main.go -env dev 使用dev配置文件
	env := flag.String("env", "dev", "setting env [dev,sit,uat,pro]")
	//自定义端口
	EnvPort = flag.String("port", ":8001", "setting env port,default 8001")

	flag.Parse()
	Yaml.SetConfigName(*env)
	//设置配置文件目录
	Yaml.AddConfigPath("./configs/yaml/")
	//设置配置文件类型
	Yaml.SetConfigType("yaml")
	if err := Yaml.ReadInConfig(); err != nil {
		panic(err)
	}
	return Yaml
}
