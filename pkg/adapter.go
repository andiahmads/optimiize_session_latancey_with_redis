package pkg

import (
	"database/sql"
	"github.com/redis/go-redis/v9"
	"gored/infra"
)

func Dbx() *sql.DB {
	return infra.GlobalConn
}

func Rdb() *redis.Client {
	return infra.DB0
}
