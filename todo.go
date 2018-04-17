package main

import (
	"database/sql"

	"github.com/t-0-m-1-3/go-echo-vue/handlers"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	// "github.com/labstack/echo/engine/standard"
)

func main() {
	db := initDB("storage.db")
	migrate(db)

	e := echo.New()

	// create the tasks routes
	e.File("/", "public/index.html")
	e.GET("/tasks", handlers.GetTasks(db))
	e.PUT("/tasks", handlers.PutTask(db))
	e.DELETE("/tasks/:id", handlers.DeleteTask(db))

	// start the server
	e.Logger.Fatal(e.Start(":8000"))
}

func initDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)

	// check for any db errors and exit
	if err != nil {
		panic(err)
	}

	// if the db connection fails without errors exit
	if db == nil {
		panic("db nil")
	}
	return db
}

func migrate(db *sql.DB) {
	sql := `
	CREATE TABLE IF NOT EXISTS tasks(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL);
		`

	_, err := db.Exec(sql)
	// exit if something goes wrong with the SQL statement
	if err != nil {
		panic(err)
	}
}
