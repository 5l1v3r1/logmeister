package logmeister

import "testing"

const TestAction = "BREW"
const TestTarget = "Teavana"
const TestResult = "418"
const TestString = "Accept-Additions: Whiskey, Cream"

func TestNewEvent(t *testing.T) {
	test_data := TestString
	e, err := NewEvent(TestAction, TestTarget, TestResult, test_data)
	if err != nil {
		t.Fatalf("Error creating new event: %v", err)
	}
	if e.Action != TestAction {
		t.Errorf("Action improperly set: %v", e.Action)
	}
	if e.Target != TestTarget {
		t.Errorf("Target improperly set: %v", e.Target)
	}
	if e.Result != TestResult {
		t.Errorf("Result improperly set: %v", e.Result)
	}
	if e.Data != TestString {
		t.Errorf("Data improperly set: %v", e.Data)
	}
}

func TestAllEmptyEvent(t *testing.T) {
	e, err := NewEvent("", "", "", "")
	EmptyTestHelper(t, e, err, EmptyActionError)
}

func TestEmptyAction(t *testing.T) {
	e, err := NewEvent("", "asdf", "asdf", "asdf")
	EmptyTestHelper(t, e, err, EmptyActionError)
}

func TestEmptyTarget(t *testing.T) {
	e, err := NewEvent("test", "", "asdf", "asdf")
	EmptyTestHelper(t, e, err, EmptyTargetError)
}

func TestEmptyResult(t *testing.T) {
	e, err := NewEvent("test", "asdf", "", "asdf")
	EmptyTestHelper(t, e, err, EmptyResultError)
}

func TestEmptyData(t *testing.T) {
	_, err := NewEvent("test", "asdf", "asdf", "")
	if err != nil {
		t.Errorf("Should not be an error, instead got: %#v", err)
	}
}

func EmptyTestHelper(t *testing.T, e *Event, err error, error_string string) {
	if e != nil {
		t.Errorf("e2: Should have returned nil.")
	}
	if err == nil {
		t.Errorf("Should have recieved error.")
	}
	if err != nil {
		if err.Error() != error_string {
			t.Errorf("Should have recieved %s, instead recieved: %v", error_string, err)
		}
	}
}
