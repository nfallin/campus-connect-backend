package Model

import (
	"image"
	"time"
)

type Campus struct {
	clubs       map[int]*Club
	users       map[int]*User
	nextEventID int
	nextClubID  int
	nextUserID  int
}

func NewCampus() *Campus {
	return &Campus{
		clubs:       make(map[int]*Club),
		users:       make(map[int]*User),
		nextEventID: 1,
		nextClubID:  1,
		nextUserID:  1,
	}
}

func (c *Campus) getNextEventID() int {
	id := c.nextEventID
	c.nextEventID++
	return id
}

func (c *Campus) getNextClubID() int {
	id := c.nextClubID
	c.nextClubID++
	return id
}

func (c *Campus) getNextUserID() int {
	id := c.nextUserID
	c.nextUserID++
	return id
}

func (c *Campus) CreateUser(name string) *User {
	newUser := &User{
		name:     name,
		id:       c.getNextUserID(),
		rsvpList: make(map[int]struct{}),
	}

	c.AddUser(newUser)
	return newUser
}

func (c *Campus) CreateClub(name string) *Club {
	newClub := &Club{
		name:   name,
		id:     c.getNextClubID(),
		events: make(map[int]*Event),
	}

	c.AddClub(newClub)
	return newClub
}

func (c *Campus) CreateEvent(name string, dateTime time.Time, location string, poster image.Image, club *Club) *Event {
	newEvent := &Event{
		name:      name,
		id:        c.getNextEventID(),
		dateTime:  dateTime,
		location:  location,
		poster:    poster,
		attendees: map[int]struct{}{},
		clubID:    club.id,
	}

	club.AddEvent(newEvent)
	return newEvent
}

func (c *Campus) DeleteEvent(e *Event, club *Club) {
	for userID := range e.attendees {
		c.users[userID].RsvpRemove(e)
	}

	club.RemoveEvent(e)
}

func (c Campus) GetClubs() map[int]*Club {
	return c.clubs
}

func (c *Campus) AddClub(club *Club) {
	c.clubs[club.id] = club
}

func (c *Campus) RemoveClub(club *Club) {
	// clean up all user's RSVP lists for all events in the club before removing the club
	for _, event := range club.GetEvents() {
		c.DeleteEvent(event, club)
	}
	delete(c.clubs, club.id)
}

func (c Campus) GetUsers() map[int]*User {
	return c.users
}

func (c *Campus) AddUser(user *User) {
	c.users[user.id] = user
}

func (c *Campus) RemoveUser(user *User) {
	// clean up all events that the user has RSVP'd to before removing the user
	for eventID := range user.rsvpList {
		if event := c.GetEvent(eventID); event != nil {
			user.RsvpRemove(event)
		}
	}
	delete(c.users, user.id)
}

func (c Campus) GetEvents() map[int]*Event {
	allEvents := make(map[int]*Event)
	for _, club := range c.clubs {
		for _, event := range club.GetEvents() {
			allEvents[event.id] = event
		}
	}
	return allEvents
}

func (c Campus) GetEvent(eventID int) *Event {
	for _, club := range c.clubs {
		if event, exists := club.GetEvents()[eventID]; exists {
			return event
		}
	}
	return nil
}
