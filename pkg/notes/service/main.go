package service

import (
	mgo "gopkg.in/mgo.v2"
)

/**
* the basic note service.
 */
type NoteService struct {
	Collection *mgo.Collection
}
