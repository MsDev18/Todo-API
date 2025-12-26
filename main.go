package main

import (
	"fmt"
	"todo-api/config"
	"todo-api/repository/mysql"
)

func main() {
	// TODO - fixed bug in declare parent name in .env file
	// for this strcure REDIS_DB_HOST
	// return redis.db.host❌ 
	// but shulde be return redisdb.host or redis-db.host ✅
	cfg := config.Load()
	fmt.Println(cfg)

	mysqlDB := mysql.New(cfg.DB)
	fmt.Println(mysqlDB)

}
