package controller

import (
	"courseCRUD/model"
	usecases_course "courseCRUD/use-cases/course"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	err := usecases_course.DeleteCourse(id)

	if err != nil {
		c.JSON(err.(model.FormateError).Code, err)
		return
	}

	c.Status(http.StatusNoContent)
}
