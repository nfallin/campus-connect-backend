package Model

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

func (c *Campus) createUser(name string) *User {
	newUser := &User{
		name:     name,
		id:       c.getNextUserID(),
		rsvpList: make(map[int]struct{}),
	}

	c.addUser(newUser)
	return newUser
}

func (c *Campus) createClub(name string) *Club {
	newClub := &Club{
		name:   name,
		id:     c.getNextClubID(),
		events: make(map[int]*Event),
	}

	c.addClub(newClub)
	return newClub
}

func (c Campus) getClubs() map[int]*Club {
	return c.clubs
}

func (c *Campus) addClub(club *Club) {
	c.clubs[club.id] = club
}

func (c *Campus) removeClub(club *Club) {
	delete(c.clubs, club.id)
}

func (c Campus) getUsers() map[int]*User {
	return c.users
}

func (c *Campus) addUser(user *User) {
	c.users[user.id] = user
}

func (c *Campus) removeUser(user *User) {
	delete(c.users, user.id)
}

// default overload: gets every event
func (c Campus) getEvents() map[int]*Event {
	allEvents := make(map[int]*Event)
	for _, club := range c.clubs {
		for _, event := range club.getEvents() {
			allEvents[event.id] = event
		}
	}
	return allEvents
}
