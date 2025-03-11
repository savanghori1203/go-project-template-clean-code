package controller

import (
	"courseCRUD/model"
	usecases_course "courseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourses(c *gin.Context) {
	response, err := usecases_course.GetCourses()

	if err != nil {
		c.IndentedJSON(err.(model.FormateError).Code, err)
		return
	}

	if len(response) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "No courses Added",
			"name":    "ValidationError",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"count":   len(response),
		"details": response,
	})
}
