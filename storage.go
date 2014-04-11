package main

import (
	"labix.org/v2/mgo"
	//"labix.org/v2/mgo/bson"
)

type Store struct {
	session *mgo.Session
	DB      *mgo.Database
}

func NewStore(url string, database string) (s *Store, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	s = &Store{session, session.DB(database)}
	return s, nil
}

func (s *Store) Close() {
	s.session.Close()
}

//func (s *Store)

// Simple wrappers to allow possible change of DBMS.
func (s *Store) Insert(collection string, docs ...interface{}) (err error) {
	return s.DB.C(collection).Insert(docs)
}

func (s *Store) Upsert(collection string, selector interface{}, update interface{}) (info *mgo.ChangeInfo, err error) {
	return s.DB.C(collection).Upsert(selector, update)
}

func (s *Store) DropCollection(collection string) (err error) {
	return s.DB.C(collection).DropCollection()
}
