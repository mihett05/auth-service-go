package services

import (
	"github.com/mihett05/auth-service-go/libs"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var _db *gorm.DB

func GetDB() *gorm.DB {
	if _db == nil {
		config := &gorm.Config{
			SkipDefaultTransaction:                   false,
			NamingStrategy:                           nil,
			Logger:                                   nil,
			NowFunc:                                  nil,
			DryRun:                                   false,
			PrepareStmt:                              false,
			DisableAutomaticPing:                     false,
			DisableForeignKeyConstraintWhenMigrating: false,
			ClauseBuilders:                           nil,
			ConnPool:                                 nil,
			Dialector:                                nil,
			Plugins:                                  nil,
		}
		db, err := gorm.Open(postgres.Open(libs.EnvDefault("DATABASE_URL")), config)
		if err != nil {
			panic(err)
		}
		_db = db
	}
	return _db
}
