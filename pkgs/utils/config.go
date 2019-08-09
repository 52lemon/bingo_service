package utils

import (
	"github.com/Unknwon/goconfig"
)

// 加载配置文件
func LoadConfigFile(sectionName string)  (config map[string]string,err error){

	// 获取配置文件路径
	cfg,err:=goconfig.LoadConfigFile("configs/run.ini")
	env, _ := cfg.GetValue("", "env")

	// 加载运行时配置
	configFile := "configs/"+env+"/config.ini"
	cfg,err=goconfig.LoadConfigFile(configFile)
	if err !=nil{
		return ;
	}

	// 获取分区配置
	config, err = cfg.GetSection(sectionName)
	if err !=nil{
		return ;
	}

	return
}