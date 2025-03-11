package main

import (
	"courseCRUD/controller"

	"github.com/gin-gonic/gin"
)

func router() *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.Greet)

	r.GET("/courses", controller.GetCourses)

	r.GET("/course/:id", controller.GetCourse)

	r.POST("/course", controller.AddCourse)

	r.PATCH("/course/:id", controller.UpdateCourse)

	r.DELETE("/course/:id", controller.DeleteCourse)

	return r
}
