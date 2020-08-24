package model

import (
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

var models []Model

func migrateAdd(model Model) {
	models = append(models, model)
}

//Model ...
type Model interface {
	Seed(db *pg.DB) error
}

//Up ...
func Up(db *pg.DB) error {
	for _, mod := range models {
		err := db.CreateTable(mod, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func Down(db *pg.DB) error {
	for _, mod := range models {
		err := db.DropTable(mod, &orm.DropTableOptions{
			IfExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func Seed(db *pg.DB) error {
	for _, mod := range models {
		err := mod.Seed(db)
		if err != nil {
			return err
		}
	}
	return nil
}
