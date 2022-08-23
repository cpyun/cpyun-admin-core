package internal

import (
	"github.com/cpyun/cpyun-admin-core/config"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/cpyun/cpyun-admin-core/database"
)

var (
	Enforcer *casbin.Enforcer
)

func NewCasbin() error {
	casbinCfg := config.Settings.Casbin
	db := database.NewBoot()

	a, err := gormadapter.NewAdapterByDB(db)
	if err != nil {
		//log.NewHelper(&log.Logger{})
		//log.NewHelper()
	}

	Enforcer, err = casbin.NewEnforcer(casbinCfg.ModelPath, a)
	if err != nil {
		return err
	}

	Enforcer.LoadPolicy()
	Enforcer.EnableLog(true)

	return nil
}
