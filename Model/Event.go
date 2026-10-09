package Model

import (
	"image"
	"time"
)

type Event struct {
	name      string
	id        int
	dateTime  time.Time
	location  string
	poster    image.Image
	attendees map[int]struct{}
	clubID    int
}

func (e *Event) GetDateTime() time.Time {
	return e.dateTime
}

func (e *Event) GetLocation() string {
	return e.location
}

func (e *Event) GetPoster() image.Image {
	return e.poster
}

func (e *Event) GetAttendees() map[int]struct{} {
	return e.attendees
}

func (e *Event) GetNumAttendees() int {
	return len(e.attendees)
}

func (e *Event) GetName() string {
	return e.name
}
