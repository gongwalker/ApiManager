package bootstrap

import (
	"ApiManager/app/Validators"
	"ApiManager/app/global"
	"database/sql"
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB

func init() {
	initConfig()
	initValidation()
	initDatabase()
}

// 加载数据库连接
func initDatabase() {
	var err error
	host := global.DbConfig["host"]
	username, _ := global.DbConfig["username"]
	password, _ := global.DbConfig["password"]
	database, _ := global.DbConfig["database"]
	port, _ := global.DbConfig["port"]
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci&loc=Local"
	DbCon, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("[MYSQL ERROR] dns" + dns + "err:" + err.Error())
	}

	err = DbCon.Ping()
	if err != nil {
		log.Fatal("[MYSQL ERROR] ", err.Error())
	}

	// 优化连接池设置
	maxIdleConns, _ := strconv.Atoi(global.DbConfig["max_idle_conns"])
	if maxIdleConns == 0 {
		maxIdleConns = 10
	}

	maxOpenConns, _ := strconv.Atoi(global.DbConfig["max_open_conns"])
	if maxOpenConns == 0 {
		maxOpenConns = 50
	}

	connMaxLifetime, _ := strconv.Atoi(global.DbConfig["conn_max_lifetime"])
	if connMaxLifetime == 0 {
		connMaxLifetime = 3600 // 默认1小时
	}

	DbCon.SetMaxIdleConns(maxIdleConns)
	DbCon.SetMaxOpenConns(maxOpenConns)
	DbCon.SetConnMaxLifetime(time.Duration(connMaxLifetime) * time.Second)
}

// 加载自定义表单验证器
func initValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("param_name", Validators.ParamName)
		_ = v.RegisterValidation("param_type", Validators.ParamType)
	}
}

// 读配置文件
func initConfig() {
	global.ReadConfig()
}
