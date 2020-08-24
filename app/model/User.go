package model

import (
	"time"

	"github.com/go-pg/pg/v9"
)

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" pg:",soft_delete"`
}

func (u *User) Seed(db *pg.DB) error {
	return nil
}

func init() {
	migrateAdd((*User)(nil))
}
