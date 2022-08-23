package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/cpyun/cpyun-admin-core/config"
	. "github.com/cpyun/cpyun-admin-core/database/logger"
	log "github.com/cpyun/cpyun-admin-core/logger"
	"github.com/cpyun/cpyun-admin-core/sdk"
)

func NewBoot() *gorm.DB {
	return open()
}

// 创建连接
func open() *gorm.DB {
	cfgMysql := config.Settings.Mysql

	mysqlConfig := mysql.Config{
		DSN:                       cfgMysql.GetMysqlDsn(), // DSN data source name
		DefaultStringSize:         191,                    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                  // 根据版本自动配置
	}
	//log.Print(cfgMysql.GetMysqlDsn())

	db, err := gorm.Open(mysql.New(mysqlConfig), getGormOption(cfgMysql.LogMode))
	if err != nil {
		panic("failed to connect database:" + err.Error())
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(cfgMysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfgMysql.MaxOpenConns)

	//db = db
	return db
}

// 获取Gorm参数
func getGormOption(mod string) *gorm.Config {
	fmt.Printf("%+v \n", log.DefaultLogger.Options())

	var cfg = &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "sd_",
			SingularTable: true,
			//NoLowerCase: true,
			//NameReplacer: nil,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		//SkipDefaultTransaction: true,		//跳过默认事务
		//Logger: nil,
		Logger: New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm(),
				),
			},
		),
	}

	return cfg
}

// 初始化数据
func Setup() {
	host := "*"
	db := open()

	sdk.Runtime.SetDb(host, db)
}
