package validations

import "db-study/pkg"

func CategoryCreateValidation(name string) error {
	if name == "" {
		return pkg.ErrorAddition(pkg.ErrValidation, "name cannot be empty")
	}

	if len(name) > 255 {
		return pkg.ErrorAddition(pkg.ErrValidation, "name is too lengthy")
	}

	return nil
}
