package initialize

import (
	"fmt"
	"gitee.com/chensyf/container/global"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func SetupMysql() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)
	dbName = global.CONFIG.Mysql.Name
	user = global.CONFIG.Mysql.User
	password = global.CONFIG.Mysql.Password
	host = global.CONFIG.Mysql.Host
	tablePrefix = global.CONFIG.Mysql.TablePrefix
	global.DB, err = gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName))
	if err != nil {
		log.Fatalf("load mysql error: %v", err)
	}
	// 处理表名
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	global.DB.SingularTable(true)
	global.DB.LogMode(true)
	global.DB.DB().SetMaxOpenConns(global.CONFIG.Mysql.MaxIdleConn)
	global.DB.DB().SetMaxOpenConns(global.CONFIG.Mysql.MaxOpenConn)
}
