package controllers

import (
	"net/http"
	"strconv"

	"github.com/Megidy/To-Do-List-Api/pkj/models"
	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var NewTask models.Task
	err := c.ShouldBindJSON(&NewTask)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Bad Request",
		})
		return
	}
	response, err := models.CreateTask(&NewTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt create new task",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"created task:":    response.Title,
		"with description": response.Description,
	})
}
func GetAllTasks(c *gin.Context) {
	tasks, err := models.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt retrieve data from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Tasks": tasks,
	})
}
func GetTaskById(c *gin.Context) {
	id := c.Param("taskId")

	taskId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt converte to int ",
		})
		return
	}
	task, err := models.GetTaskById(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt retrieve data from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{

		"task:": task,
	})

}
func DeleteTask(c *gin.Context) {
	id := c.Param("taskId")
	taskId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt converte to int ",
		})
		return
	}
	task, err := models.DeleteTask(taskId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt delete data from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Deleted Task": task,
	})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("taskId")
	taskId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt converte to int ",
		})
	}

	var UpdatedTask models.Task
	err = c.ShouldBindJSON(&UpdatedTask)

	UpdatedTask.Id = taskId
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "didnt converte to int ",
		})
		return
	}

	task, err := models.UpdateTask(&UpdatedTask)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "didnt update task ",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Updated task": task,
	})

}
