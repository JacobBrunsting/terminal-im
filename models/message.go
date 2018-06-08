package models

import (
	"time"
)

type Message struct {
	Timestamp time.Time `json:timestamp"`
	Author    Author    `json:author"`
	Text      string    `json:text"`
}
