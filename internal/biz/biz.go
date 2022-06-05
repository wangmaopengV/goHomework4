package biz

import (
	"database/sql"
)

type App struct { // 最终需要的对象
	db *sql.DB
}

var app *App

func NewApp(db *sql.DB) *App {
	app = &App{db: db}
	return app
}

func GetApp() *App {

	return app
}

func (a *App) Close() error {

	return a.db.Close()
}

func (a *App) Test() string {

	return "test"
}
