package main

import (
	"github.com/go-sql-driver/mysql"
	"github.com/yann-fk-21/expense_tracker/cmd/api"
	"github.com/yann-fk-21/expense_tracker/config"
	"github.com/yann-fk-21/expense_tracker/db"
	"github.com/yann-fk-21/expense_tracker/logger"
)

func main() {

	errLogger := logger.InitLogger()
	cfg := config.InitConfig()

	mysqlDB, err := db.NewMysqlStorage(mysql.Config{
		User:                 cfg.DBUser,
		Passwd:               cfg.DBPass,
		DBName:               cfg.DBName,
		Addr:                 cfg.DBAddr,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		errLogger.Println(err)
		errLogger.Fatal(err)
	}

	server := api.NewAPIServer(":8080", mysqlDB)
	err = server.Run()

	if err != nil {
		errLogger.Println(err)
		errLogger.Fatal(err)
	}
}
