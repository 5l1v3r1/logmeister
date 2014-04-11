package daemon

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type Store struct {
	Session *mgo.Session
	DB      *mgo.Database
}

// Store handles
func NewStore(url string, database string) (s *Store, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}
	s = &Store{session, session.DB(database)}
	return s, nil
}

func (s *Store) Close() {
	s.Session.Close()
}

// Event Store Functions
func (s *Store) StoreEvent(e *Event) (err error) {
	return s.Insert(EventCollection, e)
}

// Server Store Functions
func (s *Store) StoreServer(server *Server) (err error) {
	return s.Insert(ServerCollection, server)
}

func (s *Store) UpdateServer(server *Server) (info *mgo.ChangeInfo, err error) {
	selector := bson.M{"IP": server.IP, "Name": server.Name}
	return s.Upsert(ServerCollection, selector, server)
}

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
