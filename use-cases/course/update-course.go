package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
	"courseCRUD/exceptions"
	"courseCRUD/model"

	"github.com/google/uuid"
)

func UpdateCourse(id string, updatedCourse model.Course) error {
	_, err := uuid.Parse(id)

	if err != nil {
		return exceptions.ValidationError("EX-0001", err.Error())
	}

	if updatedCourse.IsEmpty() {
		return exceptions.ValidationError("EX-0001", "Please Enter Course Details")
	} else {
		courseDetails, err := GetCourse(id)

		if err != nil {
			return err
		}

		if len(courseDetails) == 0 {
			return exceptions.ObjectNotFoundError("EX-0004", "No Course Details were found on given ID")
		}

		return db.UpdateCourse(context.TODO(), id, updatedCourse)
	}
}
