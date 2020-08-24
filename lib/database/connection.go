package database

import (
	"github.com/go-pg/pg/v9"
	"github.com/spf13/viper"
)

var db *pg.DB

func connectionDB() *pg.DB {
	db = pg.Connect(&pg.Options{
		Addr:     viper.GetString("DB_HOST"),
		Database: viper.GetString("DB_DATABASE"),
		User:     viper.GetString("DB_USER"),
		Password: viper.GetString("DB_PASSWORD"),
	})
	return db
}

// DB get database connection
func DB() *pg.DB {
	if db == nil {
		db = connectionDB()
	}
	return db
}
