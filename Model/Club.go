package Model

type Club struct {
	events map[int]*Event
	name   string
	id     int
}

func (c *Club) getEvents() map[int]*Event {
	return c.events
}

func (c *Club) addEvent(event *Event) {
	c.events[event.id] = event
}

func (c *Club) removeEvent(event *Event) {
	delete(c.events, event.id)
}
