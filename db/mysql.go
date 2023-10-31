package db

import (
	"fmt"

	"github.com/amillerrr/htmx-todo-app/config"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

// CreateMySqlConnection creates a connection to the database
func CreateMySqlConnection(DB_NAME string) *sqlx.DB {
	env := config.LoadEnv()
	// initialize some variables for the MySQL data source
	var (
		databaseUser     string = env.DB_USER
		databasePassword string = env.DB_PASSWORD
		databaseHost     string = env.DB_HOST
		databasePort     string = env.DB_PORT
		databaseName     string = DB_NAME
	)
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	db := sqlx.MustOpen("mysql", dataSource)

	err := db.Ping()
	if err != nil {
		panic("Could not ping db")
	} else {
		println("Connected to the database")
	}

	return db
}
