package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
	"courseCRUD/exceptions"

	"github.com/google/uuid"
)

func DeleteCourse(id string) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return exceptions.ValidationError("EX-0001", err.Error())
	}

	courseDetails, err := GetCourse(id)

	if err != nil {
		return err
	}

	if len(courseDetails) == 0 {
		return exceptions.ObjectNotFoundError("EX-0004", "No Course was found on given Id")
	}

	return db.DeleteCourse(context.TODO(), id)
}
