package daemon

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

const TestCollection = "test_collection"

type TestData struct {
	ID    string
	Data  string
	Extra interface{}
}

func NewStoreHelper(t *testing.T) (s *Store) {
	s, err := NewStore("localhost", "testing")

	if err != nil {
		t.Fatalf("Error creating a new store: %v", err)
	}
	if s == nil {
		t.Fatalf("Error creating new store: nil pointer returned with no error.")
	}
	return s
}

func TestNewStore(t *testing.T) {
	s := NewStoreHelper(t)
	s.Close()
	return
}

func TestStoreEvent(t *testing.T) {
	s := NewStoreHelper(t)
	defer s.Close()

	e, err := NewEvent("Test", "Test", "Tested", "Lentils")
	if err != nil {
		t.Fatalf("Error creating new event: %v", err)
	}
	err = s.StoreEvent(e)
	if err != nil {
		t.Errorf("Error storing event: %v", err)
	}
}

func TestNewStoreBadServerUrl(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping Server URL Timeout Test to save time.")
	}
	s, err := NewStore("Vogons", "testing")
	if err == nil {
		t.Fatalf("No error. Should have returned an error or timeout.")
	}
	if s != nil {
		t.Errorf("Should return nil object. Returned: %v")
	}
	return
}

func TestInsert(t *testing.T) {
	s := NewStoreHelper(t)
	defer s.Close()

	d := &TestData{"1", "test1", "InsertTest"}
	err := s.Insert(TestCollection, d)
	if err != nil {
		t.Errorf("Error inserting into collection: %v", err)
	}
	return
}

func TestUpsert(t *testing.T) {
	s := NewStoreHelper(t)
	defer s.Close()

	selector := bson.M{"ID": "1"}
	update := &TestData{"2", "test2", "UpsertTest"}
	_, err := s.Upsert(TestCollection, selector, update)
	if err != nil {
		t.Errorf("Error upserting into collection with selector %v and update data %v: %v", selector, update, err)
	}
	return
}

func TestDropCollection(t *testing.T) {
	s := NewStoreHelper(t)
	defer s.Close()
	err := s.DropCollection(TestCollection)
	if err != nil {
		t.Errorf("Error dropping collection %s: %v", TestCollection, err)
	}
	return
}
