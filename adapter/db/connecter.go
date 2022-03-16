package db

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Dsn string

type Connector interface {
	RunInSession(f func(conn *gorm.DB) error) error
	RunInTransactionSession(f func(conn *gorm.DB) error) error
}

func NewConnector(dsn Dsn, nowf func() time.Time) Connector {
	return &accessor{
		dsn:  string(dsn),
		nowf: nowf,
	}
}

type accessor struct {
	dsn  string
	nowf func() time.Time
}

func (d *accessor) RunInSession(f func(conn *gorm.DB) error) error {
	conn, err := gorm.Open(mysql.Open(d.dsn), &gorm.Config{
		DisableAutomaticPing: true,
	})
	if err != nil {
		return err
	}
	db, err := conn.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	conn.NowFunc = d.nowf
	conn = conn.Omit(clause.Associations)
	if debug, _ := strconv.ParseBool(os.Getenv("DB_DEBUG")); debug {
		conn = conn.Debug()
	}

	return f(conn.Session(&gorm.Session{}))
}

func (d *accessor) RunInTransaction(conn *gorm.DB, f func(tx *gorm.DB) error) error {
	conn = conn.Begin()
	if conn.Error != nil {
		return conn.Error
	}
	defer func() {
		if r := recover(); r != nil {
			conn.Rollback()
			panic(r)
		}
	}()
	if err := f(conn); err != nil {
		if rberr := conn.Rollback().Error; rberr != nil {
			return fmt.Errorf("database rollback error: %s, cause: %w", rberr.Error(), err)
		}
		return err
	}
	return conn.Commit().Error
}

func (d *accessor) RunInTransactionSession(f func(conn *gorm.DB) error) error {
	return d.RunInSession(func(conn *gorm.DB) error {
		return d.RunInTransaction(conn, f)
	})
}
