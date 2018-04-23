package handlers

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/t-0-m-1-3/go-echo-vue/models"

	"github.com/labstack/echo"
)

// H maps json to an interface
type H map[string]interface{}

// GetTasks endpoint
func GetTasks(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// fetch any new tasks
		return c.JSON(http.StatusOK, models.GetTasks(db))
	}
}

// PutTasks endpoint
func PutTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// instantiate a new task
		var task models.Task
		// map incoming JSON body to the new tasks
		c.Bind(&task)
		// add a task using new model
		id, err := models.PutTask(db, task.Name)
		// Return a JSON reponse if successful
		if err != nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			//handle any errors
		} else {
			return err
		}
	}
}

// DeleteTask endpoint
func DeleteTask(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		// use out new model to delete a task
		_, err := models.DeleteTask(db, id)
		// return a json response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// handle any errors
		} else {
			return err
		}
	}
}
