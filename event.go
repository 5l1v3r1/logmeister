package logmeister

import (
	"errors"
)

const EmptyActionError = "Action must not be empty."
const EmptyTargetError = "Target must not be empty."
const EmptyResultError = "Result must not be empty."

type Event struct {
	Action string
	Target string
	Result string
	Data   string
}

// Creates a new event with the given parameters and makes sure no required fields are blank.
// General usage: Set the action against a target with given result. Often storing extra data as json.
func NewEvent(action, target, result, data string) (e *Event, err error) {
	if action == "" {
		err = errors.New(EmptyActionError)
	} else if target == "" {
		err = errors.New(EmptyTargetError)
	} else if result == "" {
		err = errors.New(EmptyResultError)
	} else {
		e = &Event{action, target, result, data}
	}
	return
}
