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
func NewStore(url string, database string) (*Store, error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	s := &Store{session, session.DB(database)}
	return s, nil
}

func (s *Store) Close() {
	s.Session.Close()
}

// Event Store Functions
func (s *Store) StoreEvent(e *Event) error {
	err := s.Insert(EventCollection, e)
	return errgo.Mask(err)
}

// Last10Events returns the last 10 events recorded in the store. If less than 10 are found,
// the size of the slice returned will be the amount actually found.
func (s *Store) Last10Events() ([]Event, error) {
	c := s.DB.C(EventCollection)
	iter := c.Find(nil).Sort("-_id").Limit(10).Iter() // Return last 10 events.
	events := []Event{}
	err := iter.All(&events)
	if err != nil {
		return nil, errgo.Mask(err)
	}
	return events, nil
}

// Simple wrappers to allow possible change of DBMS.
func (s *Store) Insert(collection string, docs ...interface{}) error {
	for _, doc := range docs {
		err := s.DB.C(collection).Insert(doc)
		if err != nil {
			return errgo.Mask(err)
		}
	}
	return nil
}

func (s *Store) Upsert(collection string, selector interface{}, update interface{}) (*mgo.ChangeInfo, error) {
	info, err := s.DB.C(collection).Upsert(selector, update)
	return info, errgo.Mask(err)
}

func (s *Store) DropCollection(collection string) error {
	err := s.DB.C(collection).DropCollection()
	return errgo.Mask(err)
}
