package infra

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gored/commons/logger"
	"log/slog"
	"os"
	"time"
)

var GlobalConn *sql.DB

func NewMySQLConn() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		logger.Slogger().Error("[INFRASTRUCTURE]", slog.Any("error:", err.Error()))
		return nil, err
	}

	GlobalConn = db
	logger.Slogger().Info("[INFRASTRUCTURE] Connected to the MySQL database!", slog.Any("msg:", db.Ping()))
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(1 * time.Hour)
	return db, nil
}
