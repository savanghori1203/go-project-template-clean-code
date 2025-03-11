package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
	"courseCRUD/exceptions"
	"courseCRUD/model"

	"github.com/google/uuid"
)

func GetCourse(id string) ([]model.Course, error) {

	_, err := uuid.Parse(id)

	if err != nil {
		return nil, exceptions.ValidationError("EX-0001", err.Error())
	}

	return db.GetCourse(context.TODO(), id)
}
