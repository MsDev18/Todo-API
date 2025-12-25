package main

import (
	"fmt"
	"todo-api/repository/mysql"
)

func main() {
	// Setup Database
	DBCfg := mysql.Config{
		DBName:   "todo_db",
		Port:     3309,
		Host:     "localhost",
		Username: "todo",
		Password: "todo-dgebwgkthv",
	}
	mysqlDB := mysql.New(DBCfg)
	fmt.Println(mysqlDB)

	
}
