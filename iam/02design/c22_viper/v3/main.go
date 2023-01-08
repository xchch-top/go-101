package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gitlab.xchch.top/golang-group/go-101/iam/02design/c22_viper/model"
	"log"
)

func main() {
	cfg := pflag.StringP("config", "c", "", "Configuration file.")
	pflag.Parse()

	if *cfg != "" {
		viper.SetConfigFile(*cfg)   // 指定配置文件名
		viper.SetConfigType("yaml") // 如果配置文件名中没有文件扩展名，则需要指定配置文件的格式，告诉viper以何种格式解析文件
	}
	// 读取配置文件，如果指定了配置文件名，则使用指定的配置文件，否则在注册的搜索路径中搜索
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s\n", err))
	}
	log.Printf("Used configuration file is: %s", viper.ConfigFileUsed())

	var C model.Config
	if err := viper.Unmarshal(&C); err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// go run main.go --config=configs/config.json
}
