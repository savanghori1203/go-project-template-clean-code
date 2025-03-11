package exceptions

import (
	"net/http"
	"time"
)

func AlreadyExistsError(customCode string, message interface{}) Error {
	if customCode == "" {
		customCode = "EX-0002"
	}

	if message == nil {
		message = "Something went wrong"
	}

	return Error{
		Name:           "AlreadyExistsError",
		Message:        message,
		CustomCode:     customCode,
		HttpStatusCode: http.StatusBadRequest,
		Date:           time.Now(),
	}
}
