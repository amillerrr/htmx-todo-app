package db

import (
	"database/sql"
	"fmt"

	"github.com/amillerrr/htmx-todo-app/utils"
	_ "github.com/lib/pq"
)

var DB *sql.DB

// CreatePostgresConnection creates a connection to the database
func init() {
	env := utils.LoadEnv()
	// initialize some variables for the MySQL data source
	var (
		databaseUser     string = env.DB_USER
		databasePassword string = env.DB_PASSWORD
		databaseHost     string = env.DB_HOST
		databasePort     string = env.DB_PORT
		databaseName     string = env.DB_NAME
	)
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	var dataSource string = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", databaseUser, databasePassword, databaseHost, databasePort, databaseName)
	var err error
	DB, err = sql.Open("postgres", dataSource)
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the postgres database")
}
