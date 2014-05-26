package logmeister

import (
	"github.com/juju/errgo"
	"labix.org/v2/mgo"
)

type Store struct {
	Session *mgo.Session
	DB      *mgo.Database
}

// Store handles
func NewStore(url string, database string) (s *Store, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	s = &Store{session, session.DB(database)}
	return s, nil
}

func (s *Store) Close() {
	s.Session.Close()
}

// Event Store Functions
func (s *Store) StoreEvent(e *Event) (err error) {
	err = s.Insert(EventCollection, e)
	return errgo.Mask(err)
}

// Simple wrappers to allow possible change of DBMS.
func (s *Store) Insert(collection string, docs ...interface{}) (err error) {
	for _, doc := range docs {
		err = s.DB.C(collection).Insert(doc)
		if err != nil {
			return errgo.Mask(err)
		}
	}
	return nil
}

func (s *Store) Upsert(collection string, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	info, err = s.DB.C(collection).Upsert(selector, update)
	return info, errgo.Mask(err)
}

func (s *Store) DropCollection(collection string) (err error) {
	err = s.DB.C(collection).DropCollection()
	return errgo.Mask(err)
}
