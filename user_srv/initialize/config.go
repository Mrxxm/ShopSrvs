package initialize

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"shop_srvs/user_srv/global"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	debug := GetEnvInfo("SHOP")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro", configFilePrefix)
	if debug == "debug" {
		configFileName = fmt.Sprintf("%s-debug", configFilePrefix)
	}

	_, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("user_srv")
	v.SetConfigName(configFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.ServerConfig); err != nil {
		panic(err)
	}
}
