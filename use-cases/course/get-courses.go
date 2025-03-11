package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
	"courseCRUD/model"
)

func GetCourses() ([]model.Course, error) {
	return db.GetCourses(context.TODO())
}
