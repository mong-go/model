package model

import (
	"errors"
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestFieldErrorsErrorFormat(t *testing.T) {
	f := FieldErrors{
		"name": []error{
			errors.New("cannot be blank"),
			errors.New("invalid format"),
		},
		"phone": []error{
			errors.New("invalid format, eg. (212) 555-1234"),
		},
	}

	assert.Equal(t, `errors:
  - name
    * cannot be blank
    * invalid format
  - phone
    * invalid format, eg. (212) 555-1234`, f.Error())
}
