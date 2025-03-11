package controller

import (
	"courseCRUD/model"
	usecases_course "courseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")

	var updatedDetails model.Course

	if err := c.BindJSON(&updatedDetails); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"name":    "ValidationError",
		})
		return
	}

	err := usecases_course.UpdateCourse(id, updatedDetails)

	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Updated Successfully",
	})
}
