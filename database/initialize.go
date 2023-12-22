package database

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/cpyun/cpyun-admin-core/config"
	log "github.com/cpyun/cpyun-admin-core/logger"
	"github.com/cpyun/cpyun-admin-core/sdk"
	"github.com/cpyun/cpyun-admin-core/sdk/pkg"
	toolsDB "github.com/cpyun/cpyun-admin-core/tools/database"
	toolLogger "github.com/cpyun/cpyun-admin-core/tools/gorm/logger"
)

// Setup 配置数据库
func Setup() {
	for k := range config.DatabasesConfig {
		openDatabase(k, config.DatabasesConfig[k])
	}
}

// 创建连接
func openDatabase(host string, c *config.Database) {
	log.Debugf("Database at [%s] => %s", host, pkg.Green(c.Source))

	registers := make([]toolsDB.ResolverConfigure, len(c.Registers))
	for i := range c.Registers {
		registers[i] = toolsDB.NewResolverConfigure(
			c.Registers[i].Sources,
			c.Registers[i].Replicas,
			c.Registers[i].Policy,
			c.Registers[i].Tables)
	}
	resolverConfig := toolsDB.NewConfigure(c.Source, c.MaxIdleConns, c.MaxOpenConns, c.ConnMaxIdleTime, c.ConnMaxLifeTime, registers)
	db, err := resolverConfig.Init(getGormOption(c.LoggerMode), opens[c.Driver])
	if err != nil {
		log.Fatal("failed to connect database:" + err.Error())
	}

	//sqlDB, _ := db.DB()
	//defer sqlDB.Close()
	//sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	//sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	sdk.Runtime.SetDb(host, db)
	//sdk.Runtime.SetCasbin(host, e)
}

// 获取Gorm参数
func getGormOption(mod string) *gorm.Config {
	mod = ""
	return &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix: "sd_",
			SingularTable: true,
			//NoLowerCase: true,
			//NameReplacer: nil,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		//SkipDefaultTransaction: true,		//跳过默认事务
		//Logger: nil,
		Logger: toolLogger.New(
			logger.Config{
				SlowThreshold: time.Second,
				Colorful:      false,
				LogLevel: logger.LogLevel(
					log.DefaultLogger.Options().Level.LevelForGorm(),
				),
			},
		),
	}
}
