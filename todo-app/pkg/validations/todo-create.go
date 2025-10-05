package validations

import (
	"db-study/pkg"
)

func CreateTodoValidation(title string) error {

	if title == "" {
		return pkg.ErrorAddition(pkg.ErrValidation, "title cannot be empty")
	}

	if len(title) > 255 {
		return pkg.ErrorAddition(pkg.ErrValidation, "title is too lengthy")
	}

	return nil
}
