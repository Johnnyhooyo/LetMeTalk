package global

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)

// ProjectConfig 全局配置
var ProjectConfig Configuration

// Configuration 全局配置信息
type Configuration struct {
	Environment string `mapstructure:"env" json:"env" yaml:"env"` // 环境 DEV/TEST/SANDBOX/PROD
	ChatAIKey   string `mapstructure:"chat_ai_key" json:"chat_ai_key" yaml:"chat_ai_key" `
}

// Viper 全局配置
var Viper *viper.Viper

// InitViper 加载配置信息 优先级: 命令行 > 环境变量 > 默认值
func InitViper(path ...string) {
	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose config file.")
		flag.Parse()
		if config == "" {
			config = ConfigFile
			fmt.Printf("您正在使用config的默认值,config的路径为%v\n", ConfigFile)
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用func Viper()传递的值,config的路径为%v\n", config)
	}

	// config = convertIfRelativePath(config)

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatalln("Fatal error config file:", err)
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&ProjectConfig); err != nil {
			log.Fatalln("配置信息变更后拉取失败:", err.Error())
		}
	})
	if err := v.Unmarshal(&ProjectConfig); err != nil {
		log.Fatalln("配置信息拉取失败", err.Error())
	}
	Viper = v
}
