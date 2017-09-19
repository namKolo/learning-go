package model

import (
	"time"
)

type Item struct {
	ID        string    `json:"id" gorethink:"id,omitempty"`
	Text      string    `json:"text" gorethink:"text"`
	Status    string    `json:"status" gorethink:"status"`
	CreatedAt time.Time `json:"createdAt" gorethink:"createdAt"`
}

func (t *Item) Completed() bool {
	return t.Status == "complete"
}

func NewItem(text string) *Item {
	return &Item{
		Text:   text,
		Status: "active",
	}
}
