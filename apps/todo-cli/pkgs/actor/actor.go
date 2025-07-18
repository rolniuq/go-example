package actor

import "fmt"

type Action string

const (
	Add            Action = "add"
	Remove         Action = "remove"
	MarkCompleted  Action = "mark_completed"
	MarkIncomplete Action = "mark_incomplete"
)

type Actor struct{}

func NewActor() *Actor {
	return &Actor{}
}

func (a *Actor) Validate(action Action) error {
	if a == nil {
		return nil
	}

	if action != Add && action != Remove && action != MarkCompleted && action != MarkIncomplete {
		return fmt.Errorf("invalid action: %s", action)
	}

	return nil
}
