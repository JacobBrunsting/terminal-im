package models

import (
	"fmt"
)

type NotFound struct {
	Name string
}

func (e *NotFound) Error() string {
	return fmt.Sprintf("could not find room '%v'\n", e.Name)
}

func IsNotFound(err error) bool {
	_, ok := err.(*NotFound)
	return ok
}

type NameConflict struct{}

func (e *NameConflict) Error() string {
	return "room being stored has a name that already exists"
}

func IsNameConflict(err error) bool {
	_, ok := err.(*NameConflict)
	return ok
}
