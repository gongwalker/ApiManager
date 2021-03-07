package global

import (
	"ApiManager/app/libs"
	"log"
	"os"
)

var (
	DbConfig                = make(map[string]string)
	SiteConfig              = make(map[string]string)
	SessionDrive            string
	SessionOption           = make(map[string]interface{})
	SessionDriveRedisConfig = make(map[string]interface{})
)

func ReadConfig() {
	rootPath, _ := os.Getwd()
	config, err := libs.ReadIniFile(rootPath + "/config/config.ini")

	if err != nil {
		log.Fatal(err)
	}

	// 站点配置
	SiteConfig["http_port"], _ = config.GetConfig("site.http_port")

	// 数据库配置
	DbConfig["host"], _ = config.GetConfig("mysql.host")
	DbConfig["username"], _ = config.GetConfig("mysql.username")
	DbConfig["password"], _ = config.GetConfig("mysql.password")
	DbConfig["database"], _ = config.GetConfig("mysql.database")
	DbConfig["port"], _ = config.GetConfig("mysql.port")

	// session驱动类型
	SessionDrive, _ = config.GetConfig("session.driver")

	// session 配置项
	SessionOption["max_age"], _ = config.GetConfigToInt("session.max_age")
	SessionOption["path"], _ = config.GetConfig("session.path")
	SessionOption["http_only"], _ = config.GetConfigToBool("session.http_only")
	SessionOption["domain"], _ = config.GetConfig("session.domain")
	SessionOption["secure"], _ = config.GetConfigToBool("session.secure")

	// 若session驱动为redis则读redis配置项
	if SessionDrive == "redis" {
		SessionDriveRedisConfig["size"], _ = config.GetConfigToInt("session_driver_redis.max_idel_con")
		SessionDriveRedisConfig["network"], _ = config.GetConfig("session_driver_redis.network")
		SessionDriveRedisConfig["address"], _ = config.GetConfig("session_driver_redis.address")
		SessionDriveRedisConfig["password"], _ = config.GetConfig("session_driver_redis.password")
	}

}
