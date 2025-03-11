package controller

import (
	"courseCRUD/model"
	usecases_course "courseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCourse(c *gin.Context) {
	id := c.Param("id")

	response, err := usecases_course.GetCourse(id)

	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}

	if len(response) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": "No course found on given Id",
			"name":    "ValidationError",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"details": response,
	})
}
