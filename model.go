package model

import (
	"gopkg.in/mgo.v2"
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
	// error if the model is invalid. Errors are not resent on each validate call
	// that is up to userland to handle any reseting between validations on a
	// singular object
	Validate(WriteType, *mgo.Database) error
}
