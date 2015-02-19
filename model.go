package model

import (
	"gopkg.in/mgo.v2"
)

// ModelReader is an interface representing a struct used only for queriyng
type ModelReader interface {
	// Collection is the name of the mongodb collection
	Collection() string
}

// ModelWriter is an interface representing a struct used for both reading and
// writing to a mongo database
type ModelWriter interface {
	ModelReader

	Valid(*mgo.Database) error
}
