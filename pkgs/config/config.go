package config

import (
	"bingo_service/pkgs/utils"
)

// 数据库配置
func GetDbConfig()  map[string]string {

	// 初始化数据库配置map
	dbConfig,_:=utils.LoadConfigFile("mysql")
	return dbConfig
}

// 服务配置
func GetServerConfig() map[string]string {

	// 初始化服务器配置map
	serverConfig,_:=utils.LoadConfigFile("server")
	return serverConfig
}