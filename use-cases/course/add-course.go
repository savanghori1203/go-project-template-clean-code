package usecases_course

import (
	"context"
	db "courseCRUD/data-access"
	"courseCRUD/exceptions"
	"courseCRUD/model"
	"fmt"
	"strings"

	"github.com/go-playground/validator"
)

func AddCourse(name string, platform string, price string) (string, error) {
	validate := validator.New()

	course := model.Course{
		Name:     name,
		Platform: platform,
		Price:    price,
	}

	if err := validate.Struct(course); err != nil {
		errors := make(map[string]string)

		for _, error := range err.(validator.ValidationErrors) {
			fmt.Println(error)
			errors[error.Field()] = error.Tag()
		}

		var validationErrors []string

		for field, tag := range errors {
			if tag == "required" {
				validationErrors = append(validationErrors, strings.ToLower(field)+" is required")
			} else if tag == "min" {
				validationErrors = append(validationErrors, strings.ToLower(field)+" should be minimum of length 2")
			} else if tag == "number" {
				validationErrors = append(validationErrors, strings.ToLower(field)+" must be a number")
			}
		}

		return "", exceptions.ValidationError("EX-0001", validationErrors)
	}

	if course.IsEmpty() {
		return "", exceptions.ValidationError("EX-0001", "Please enter a course details")
	} else {
		isCourseExists, err := isCourseExists(name)

		if err != nil {
			return "", err
		}

		if isCourseExists {
			return "", exceptions.AlreadyExistsError("EX-0003", "Course details already exists with given Name")
		}

		return db.AddCourse(context.TODO(), name, platform, price)
	}

}
