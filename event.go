package logmeister

import (
	"errors"
	"time"
)

const EmptyActionError = "Action must not be empty."
const EmptyTargetError = "Target must not be empty."
const EmptyResultError = "Result must not be empty."

// The general data strucutre used by logmeister. This is for very basic monitoring/event logging.
type Event struct {
	Time   string // Time is often time.Now().String() as it can the be parsed easily by the time package.
	Action string // Action should be the verb of what happened. e.g. "Purchase", "Ping", "Boil"
	Target string // Target should be the Noun being acted on. For example, the server you are pinging.
	Result string // Result is usually the response you get from the action. Maybe JSON for resp.StatusCode and resp.Body?
	Data   string // Data is extra data you wish to store that is directly related to this event. Generally JSON.
}

// Creates a new event with the given parameters and makes sure no required fields are blank.
func NewEvent(action, target, result, data string) (e *Event, err error) {
	if action == "" {
		err = errors.New(EmptyActionError)
	} else if target == "" {
		err = errors.New(EmptyTargetError)
	} else if result == "" {
		err = errors.New(EmptyResultError)
	} else {
		e = &Event{time.Now().String(), action, target, result, data}
	}
	return
}
