package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"log"
)

var (
	cfg  = pflag.StringP("config", "c", "", "Configuration file.")
	help = pflag.BoolP("help", "h", false, "Show this help message.")
)

func main() {
	pflag.Parse()
	if *help {
		pflag.Usage()
		return
	}

	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	} else {
		viper.AddConfigPath(".") // 把当前目录加入到配置文件的搜索路径中，配置文件搜索路径，可以设置多个搜索路径
		// viper.AddConfigPath("$HOME/.iam")
		viper.SetConfigName("config") // 配置文件名称（没有文件扩展名）
	}

	// 读取配置文件，如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}

	log.Printf("Used configuration file is: %s", viper.ConfigFileUsed())

	// 示例1
	// go run main.go
	// panic: Fatal error config file: Config File "config" Not Found in "[D:\\gitlab-clone\\golang-group\\go-101\\iam\\02design\\c22_viper\\v1]"

	// 示例2
	// go run main.go --config=configs/config.yaml
	// Used configuration file is: configs/config.yaml
}
