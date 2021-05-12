package initialize

import (
	"flag"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"gin-vben-admin/common"
	"os"
)



func InitConf(path ...string) {

	var config string
	if len(path) == 0 {
		flag.StringVar(&config, "c", "", "choose conf file.")
		flag.Parse()
		if config == "" { // 优先级: 命令行 > 环境变量 > 默认值
			if configEnv := os.Getenv("CONFFILE"); configEnv == "" {
				config = common.ConfigFile
				fmt.Printf("您正在使用config的默认值,config的路径为%v\n", config)
			} else {
				config = configEnv
				fmt.Printf("您正在使用configEnv环境变量,config的路径为%v\n", config)
			}
		} else {
			fmt.Printf("您正在使用命令行的-c参数传递的值,config的路径为%v\n", config)
		}
	} else {
		config = path[0]
		fmt.Printf("您正在使用InitConf传递的值,config的路径为%v\n", config)
	}
	v := viper.New()
	v.SetConfigFile(config)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error conf file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("conf file changed:", e.Name)
		if err = v.Unmarshal(&common.CONFIG); err != nil {
			fmt.Println(err)
		}
	})
	if err = v.Unmarshal(&common.CONFIG); err != nil {
		fmt.Println(err)
	}
	common.VP = v
}
