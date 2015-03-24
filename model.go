package model

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// ModelReader is an interface representing a struct used only for queriyng
type ModelReader interface {
	// CollectionName is the name of the mongodb collection
	CollectionName() string
}

// WriteType is a SAVE or UPDATE
type WriteType int

const (
	_ WriteType = iota
	SAVE
	UPDATE
)

// ModelWriter is an interface representing a struct used for both reading and
// writing to a mongo database
type ModelWriter interface {
	ModelReader

	// Validate is intented to run before a save/update call and will return an
	// error if the model is invalid. Errors are not reset on each validate call
	// that is up to userland to handle any reseting between validations on a
	// singular object
	Validate(WriteType, *mgo.Database) error
}

// Model is a basic embeddable struct providing some standard fields and helper
// methods
type Model struct {
	ID bson.ObjectId `bson:"_id,omitempty" json:"_id,omitempty"`

	// FieldErrors holds fields errors that have been gathered through Validates()
	// This field should not be called directly, but must be exposed for encoding.
	FieldErrors FieldErrors `bson:"-" json:"errors,omitempty"`
}

func (m *Model) AddError(name string, err error) {
	if m.FieldErrors == nil {
		m.FieldErrors = make(FieldErrors)
	}

	m.FieldErrors[name] = append(m.FieldErrors[name], err)
}

func (m *Model) Errors() error {
	if len(m.FieldErrors) == 0 {
		return nil
	}

	return m.FieldErrors
}

func (m *Model) ResetErrors() {
	m.FieldErrors = nil
}
