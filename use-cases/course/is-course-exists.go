package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
)

func isCourseExists(name string) (bool, error) {
	courseDetails, err := db.GetCourseByName(context.TODO(), name)

	if err != nil {
		return false, err
	}

	return len(courseDetails) != 0, nil
}
