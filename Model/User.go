package Model

import (
	"image"
	"time"
)

type User struct {
	name     string
	id       int
	rsvpList map[int]struct{}
}

// returns the list of events that u has RSVP'd to
func (u User) getRsvpList() map[int]struct{} {
	return u.rsvpList
}

// add e to u's rsvpList
func (u *User) rsvpAdd(e *Event) {
	u.rsvpList[e.id] = struct{}{}
	e.attendees[u.id] = struct{}{}
}

// remove e from u's rsvpList
func (u *User) rsvpRemove(e *Event) {
	delete(u.rsvpList, e.id)
	delete(e.attendees, u.id)
}

func (u User) createEvent(campus *Campus, club *Club, name string, dateTime time.Time, location string, poster image.Image) *Event {
	newEvent := &Event{
		name:      name,
		id:        campus.getNextEventID(),
		dateTime:  dateTime,
		location:  location,
		poster:    poster,
		attendees: map[int]struct{}{},
		clubID:    club.id,
	}

	club.addEvent(newEvent)
	return newEvent
}

func (u User) deleteEvent(e *Event, club *Club) {
	club.removeEvent(e)
}
