package model

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"gopkg.in/nowk/assert.v2"
)

func TestModelEmbeddedEncoding(t *testing.T) {
	m := tModel{
		Name: "foo",
	}
	m.AddError("name", errors.New("cannot be blank"))

	w := httptest.NewRecorder()
	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, `{"errors":{"name":["cannot be blank"]},"name":"foo"}`+"\n",
		w.Body.String())
}
