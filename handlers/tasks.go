package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/t-0-m-1-3/go-echo-vue/models"

	"github.com/labstack/echo"
)

// H creates an interface for the JSON
type H map[string]interface{}

//GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch the tasks using the model
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTask endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instantiate a new task
		var task models.Task
		// map incoming JSON body to new task
		c.Bind(&task)
		// Add a task using new model
		id, err := models.PutTask(db, task.Name)
		// Return a status is success
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// handle the errors
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		//use new model to delete a task
		_, err := models.DeleteTask(db, id)
		// return a json success msg
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// handle errors
		} else {
			return err
		}
	}
}
