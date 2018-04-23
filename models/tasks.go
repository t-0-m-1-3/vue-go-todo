package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// Task is a struct containing Task data
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TaskCollection is a collection of tasks
type TaskCollection struct {
	Tasks []Task `json:"items"`
}

// GetTasks gets all the data
func GetTasks(db *sql.DB) TaskCollection {
	sql := "SELECT * FROM tasks"
	rows, err := db.Query(sql)
	// exit if sql doesn't work
	if err != nil {
		panic(err)
	}

	//make sure to clean up resources
	defer rows.Close()

	result := TaskCollection{}
	for rows.Next() {
		task := Task{}
		err2 := rows.Scan(&task.ID, &task.Name)
		// exit if we get an error
		if err2 != nil {
			panic(err2)
		}
		result.Tasks = append(result.Tasks, task)
	}
	return result
}

// PutTask inserts a new task
func PutTask(db *sql.DB, name string) (int64, error) {
	sql := "INSERT INTO tasks(name) VALUES(?)"

	// create a prepared sql statement
	stmt, err := db.Prepare(sql)
	// exit if we get an error
	if err != nil {
		panic(err)
	}

	// make sure to close resources
	defer stmt.Close()

	// replace the ? with name
	result, err2 := stmt.Exec(name)
	// exit if we get an error
	if err2 != nil {
		panic(err2)
	}

	return result.LastInsertId()
}

// DeleteTask removes a task from the DB
func DeleteTask(db *sql.DB, id int) (int64, error) {
	sql := "DELETE FROM tasks WHERE id = ?"

	// create a prepared sql statement
	stmt, err2 := db.Prepare(sql)
	// exit for error
	if err != nil {
		panic(err)
	}

	//Replace the ? with id
	result, err2 := stmt.Exec(id)
	//exit for errors
	if err2 != nil {
		panic(err2)
	}
	return result.RowsAffected()
}
