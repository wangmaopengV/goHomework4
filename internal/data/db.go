package data

import (
	"database/sql"
	"fmt"
	"goHomework4/internal/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
)

var Provider = wire.NewSet(New)

func New(cfg *config.MysqlConf) (db *sql.DB, err error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Pwd, cfg.Host, cfg.Port, cfg.DbName)

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	if err = db.Ping(); err != nil {
		return
	}
	return db, nil
}
