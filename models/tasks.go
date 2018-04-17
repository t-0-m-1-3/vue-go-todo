package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is a collection of Tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks grabs the tasks from the db and puts them in a collection
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// Exit if the SQL doesn't work
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := TaskCollection{}

	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		//Exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

// PutTask inserts tasks into the databse
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// Create a prepared sql statement
	stmt, err := db.Prepare(sql)
	// exit if we get an error
	if err != nil {
		panic(nil)
	}
	// clean up after exit
	defer stmt.Close()

	// Replace the ? with prepared statement
	result, err2 := stmt.Exec(name)
	// exit if we get an error
	if err2 != nil {
		panic(err2)
	}
	return result.LastInsertId()
}

// DeleteTask remotes a task from the DB
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// create the prepared SQL statement
	stmt, err := db.Prepare(sql)
	// Exit if we get an error
	if err != nil {
		panic(err)
	}
	// Replace the ? with the statement
	result, err2 := stmt.Exec(id)

	//Exit for errors
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
