package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var opens = map[string]func(string) gorm.Dialector{
	"mysql": mysql.Open,
	//"postgres": postgres.Open,
	//"sqlite3":  sqlite.Open,
	//"clickhouse": clickhouse.Open,
}

func WithOpens(key string, open func(s string) gorm.Dialector) {
	opens[key] = open
}
