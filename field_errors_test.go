package model

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
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

type model struct {
	Model `bson:",inline"`
}

func TestErrorsJsonEncoding(t *testing.T) {
	m := model{}
	m.AddError("name", errors.New("cannot be blank"))

	w := httptest.NewRecorder()
	err := json.NewEncoder(w).Encode(&m)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, `{"errors":{"name":["cannot be blank"]}}`+"\n",
		w.Body.String())
}
