package model

import (
	"errors"
)

type AddErrorer interface {
	AddError(string, error)
}

// Assert is a helper method to streamline validation checking and adding errors
//
// Example:
//		model.Assert(&m, m.Name == "", "name", "cannot be blank")
// Or
//		model.Assert(&m, m.Name == "", "name", errors.New("cannot be blank"))
func Assert(m AddErrorer, invalid bool, name string, e interface{}) bool {
	if invalid {
		var err error
		switch v := e.(type) {
		case string:
			err = errors.New(v)
		case error:
			err = v
		}

		m.AddError(name, err)
	}

	// negate invalid
	return !invalid
}
