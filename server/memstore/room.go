package memstore

import (
	"fmt"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/jbrunsting/terminal-im/models"
)

type NotFound struct {
    name string
}

func (e *NotFound) Error() string {
    return fmt.Sprintf("could not find room '%v'\n", e.name)
}

func IsNotFound(err error) bool {
	_, ok := err.(*NameConflict)
	return ok
}

type NameConflict struct {}

func (e *NameConflict) Error() string {
    return "room being stored has a name that already exists"
}

func IsNameConflict (err error) bool {
	_, ok := err.(*NameConflict)
	return ok
}

type RoomStore interface {
	StoreRoom(r *models.Room) error
	RetrieveRoom(name string) (*models.Room, error)
}

type roomStore struct {
	rooms *cache.Cache
}

func NewRoomStore() RoomStore {
	return &roomStore{
		rooms: cache.New(24 * time.Hour, 15 * time.Minute),
	}
}

func (s *roomStore) StoreRoom(r *models.Room) error {
	if _, ok := s.rooms.Get(r.Name); ok {
        return &NameConflict{}
	}
	s.rooms.SetDefault(r.Name, r)
	return nil
}

func (s *roomStore) RetrieveRoom(name string) (*models.Room, error) {
	if r, ok := s.rooms.Get(name); ok {
		// replace the room so the cache expiration time is updated
		s.rooms.SetDefault(name, r)
		return r.(*models.Room), nil
	}
	return nil, &NotFound{name}
}
