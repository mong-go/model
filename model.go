package model

import (
	"gopkg.in/mgo.v2"
)

// ModelReader is an interface representing a struct used only for queriyng
type ModelReader interface {
	// Collection is the name of the mongodb collection
	Collection() string
}

// WriteType is a Save or Update
type WriteType int

const (
	_ WriteType = iota
	Save
	Update
)

// ModelWriter is an interface representing a struct used for both reading and
// writing to a mongo database
type ModelWriter interface {
	ModelReader

	// Valid is intented to run before a save/update call and will return an error
	// if the model is invalid.
	Valid(WriteType, *mgo.Database) error
}
