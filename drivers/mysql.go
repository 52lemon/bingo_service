// mysql db drives
package drivers

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"
	"bingo_service/pkgs/config"
	_ "github.com/go-sql-driver/mysql"

)

// query need rows.Close to release db ins
// exec will release automatic
var MysqlDb *sql.DB  // db pool instance
var MysqlDbErr error // db err instance

func init() {

	// 获取db配置
	dbConfig := config.GetDbConfig()

	dbDSN := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		dbConfig["db_user"],
		dbConfig["db_pwd"],
		dbConfig["db_host"],
		dbConfig["db_port"],
		dbConfig["db_name"],
		dbConfig["db_charset"],
	)

	// 打开db
	MysqlDb, MysqlDbErr = sql.Open("mysql", dbDSN)

	if MysqlDbErr != nil {
		panic("database data source name error: " + MysqlDbErr.Error())
	}

	// 连接池最大连接数
	dbMaxOpenConns, _ := strconv.Atoi(dbConfig["db_max_open_conns"])
	MysqlDb.SetMaxOpenConns(dbMaxOpenConns)

	// 连接池最大空闲数
	dbMaxIdleConns, _ := strconv.Atoi(dbConfig["db_max_idle_conns"])
	MysqlDb.SetMaxIdleConns(dbMaxIdleConns)

	// 连接池链接最长生命周期
	dbMaxLifetimeConns, _ := strconv.Atoi(dbConfig["db_max_lifetime_conns"])
	MysqlDb.SetConnMaxLifetime(time.Duration(dbMaxLifetimeConns))

	if MysqlDbErr = MysqlDb.Ping(); nil != MysqlDbErr {
		panic("database connect failed: " + MysqlDbErr.Error())
	}
}
