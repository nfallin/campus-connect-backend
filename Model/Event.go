package Model

import (
	"image"
	// "image/png"
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

// events are constructed by users

func (e *Event) getDateTime() time.Time {
	return e.dateTime
}

func (e *Event) getLocation() string {
	return e.location
}

func (e *Event) getPoster() image.Image {
	return e.poster
}

func (e *Event) getAttendees() map[int]struct{} {
	return e.attendees
}

func (e *Event) getNumAttendees() int {
	return len(e.attendees)
}

func (e *Event) getName() string {
	return e.name
}
