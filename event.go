package daemon

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
