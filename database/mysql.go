package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/cpyun/cpyun-admin-core/config"
	dbLogger "github.com/cpyun/cpyun-admin-core/database/logger"
	log "github.com/cpyun/cpyun-admin-core/logger"
	"github.com/cpyun/cpyun-admin-core/sdk"
	toolsDB "github.com/cpyun/cpyun-admin-core/tools/database"
)

//
func NewBoot() *gorm.DB {
	host := "*"
	db := sdk.Runtime.GetDbByKey(host)
	return db
}

// 初始化数据
func Setup() {
	host := "*"

	db := sdk.Runtime.GetDbByKey(host)
	if db == nil {
		db = openDatabase()
		sdk.Runtime.SetDb(host, db)
	}
}

// 创建连接
func openDatabase() *gorm.DB {
	cfgMysql := config.Settings.Mysql

	//mysqlConfig := mysql.Config{
	//	DSN:                       cfgMysql.GetMysqlDsn(), // DSN data source name
	//	DefaultStringSize:         191,                    // string 类型字段的默认长度
	//	DisableDatetimePrecision:  true,                   // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	//	DontSupportRenameIndex:    true,                   // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	//	DontSupportRenameColumn:   true,                   // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	//	SkipInitializeWithVersion: false,                  // 根据版本自动配置
	//}

	registers := make([]toolsDB.ResolverConfigure, len(cfgMysql.Registers))
	for i := range cfgMysql.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			cfgMysql.Registers[i].Sources,
			cfgMysql.Registers[i].Replicas,
			cfgMysql.Registers[i].Policy,
			cfgMysql.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(cfgMysql.GetMysqlDsn(), cfgMysql.MaxIdleConns, cfgMysql.MaxOpenConns, cfgMysql.ConnMaxIdleTime, cfgMysql.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(getGormOption(cfgMysql.LogMode), opens["mysql"])
	if err != nil {
		log.Fatal("failed to connect database:" + err.Error())
	}

	sqlDB, _ := db.DB()
	//defer sqlDB.Close()
	sqlDB.SetMaxIdleConns(cfgMysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfgMysql.MaxOpenConns)

	return db
}

// 获取Gorm参数
func getGormOption(mod string) *gorm.Config {
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
		Logger: dbLogger.New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      false,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm(),
				),
			},
		),
	}

	return cfg
}
