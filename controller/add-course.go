package controller

import (
	"courseCRUD/exceptions"
	"courseCRUD/model"
	usecases_course "courseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddCourse(c *gin.Context) {
	var course model.Course

	err := c.BindJSON(&course)

	if err != nil {
		c.JSON(http.StatusBadRequest, exceptions.ValidationError("EX-0001", err.Error()))
		return
	}

	id, err := usecases_course.AddCourse(course.Name, course.Platform, course.Price)

	if err != nil {
		c.JSON(err.(exceptions.Error).HttpStatusCode, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}
