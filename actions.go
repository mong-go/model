package model

import (
	"errors"

	"gopkg.in/mgo.v2"
)

var ErrInvalidCollectionName = errors.New("invalid collection name")

// Save saves a ModelWriter to db
func Save(m ModelWriter, db *mgo.Database) error {
	err := m.Validate(SAVE, db)
	if err != nil {
		return err
	}

	c, err := Getc(m, db)
	if err != nil {
		return err
	}

	return c.Insert(m)
}

// TODO ability to save multiple ModelWriters

// Update saves a ModelWriter to db, against a given selector
func Update(sel interface{}, m ModelWriter, db *mgo.Database) error {
	err := m.Validate(UPDATE, db)
	if err != nil {
		return err
	}

	c, err := Getc(m, db)
	if err != nil {
		return err
	}

	return c.Update(sel, m)
}

// TODO UpdateAll
// TODO Remove
// TODO RemoveAll

// Getc returns the db collection from the given interface. Returns an error if
// the collection name is blank
func Getc(m ModelReader, db *mgo.Database) (*mgo.Collection, error) {
	name := m.CollectionName()
	if name == "" {
		return nil, ErrInvalidCollectionName
	}

	return db.C(name), nil
}
