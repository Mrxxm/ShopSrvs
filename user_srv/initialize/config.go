package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"shop_srvs/user_srv/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func GetEnvInfo(env string) string {
	viper.AutomaticEnv()
	return viper.GetString(env)
}

func InitConfig() {
	// 一、本地获取nacos配置
	debug := GetEnvInfo("SHOP")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro", configFilePrefix)
	if debug == "debug" {
		configFileName = fmt.Sprintf("%s-debug", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigType("yaml")
	v.AddConfigPath("user_srv")
	v.SetConfigName(configFileName)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.Unmarshal(&global.NacosConfig); err != nil {
		panic(err)
	}

	// 二、nacos服务中获取其他配置
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.Nacos.Host,
			Port:   uint64(global.NacosConfig.Nacos.Port),
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.Nacos.Namespace, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		zap.S().Info("创建nacos客户端异常：", err.Error())
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.Nacos.DataId,
		Group:  global.NacosConfig.Nacos.Group})

	if err != nil {
		zap.S().Info("获取nacos配置异常：", err.Error())
	}

	// json转换映射成结构体
	jsonBytesContent := []byte(content)
	err = json.Unmarshal(jsonBytesContent, &global.ServerConfig)
	if err != nil {
		zap.S().Info("转换nacos配置异常：", err.Error())
	}
	zap.S().Infof("转换nacos配置打印：%+v", global.ServerConfig)
}

func InitConfig2() {
	debug := GetEnvInfo("SHOP")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s-pro", configFilePrefix)
	if debug == "debug" {
		configFileName = fmt.Sprintf("%s-debug", configFilePrefix)
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
