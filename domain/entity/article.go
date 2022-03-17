package entity

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type Article struct {
	ID        string    `gorm:"primary_key;column:id"`
	Title     string    `gorm:"column:title"`
	Body      string    `gorm:"column:body"`
	UserID    UserID    `gorm:"column:user_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`

	// belongs to
	User *User `gorm:"foreignkey:UserID"`
}
