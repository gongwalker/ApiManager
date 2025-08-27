package global

import (
	"ApiManager/app/libs"
	"log"
	"os"
	"strings"
)

var (
	DbConfig                = make(map[string]string)
	SiteConfig              = make(map[string]string)
	SessionDrive            string
	SessionOption           = make(map[string]interface{})
	SessionDriveRedisConfig = make(map[string]interface{})
	GinRunMode              string
	GinWriteLog             bool
	EnableDomainCheck       bool     // 是否启用域名访问限制
	AllowedHosts            []string // 允许访问的域名列表
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
	DbConfig["max_idle_conns"], _ = config.GetConfig("mysql.max_idle_conns")
	DbConfig["max_open_conns"], _ = config.GetConfig("mysql.max_open_conns")
	DbConfig["conn_max_lifetime"], _ = config.GetConfig("mysql.conn_max_lifetime")

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

	// 运行模式
	GinRunMode, err = config.GetConfig("site.gin_run_mode")

	// 是否记录运行日志
	GinWriteLog, err = config.GetConfigToBool("site.gin_write_log")

	// 是否启用域名访问限制
	EnableDomainCheck, err = config.GetConfigToBool("site.enable_domain_check")

	// 读取允许访问的域名列表
	allowedHostsStr, _ := config.GetConfig("site.allowed_hosts")
	if allowedHostsStr != "" {
		// 将逗号分隔的域名列表转换为字符串数组
		AllowedHosts = strings.Split(allowedHostsStr, ",")
		// 去除每个域名的空白字符
		for i, host := range AllowedHosts {
			AllowedHosts[i] = strings.TrimSpace(host)
		}
	} else {
		// 默认允许的域名列表
		AllowedHosts = []string{
			"10.0.7.128:8080",
			"127.0.0.1:8080",
			"localhost:8080",
		}
	}
}
