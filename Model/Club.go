package Model

type Club struct {
	events map[int]*Event
	name   string
	id     int
}

func (c *Club) GetEvents() map[int]*Event {
	return c.events
}

func (c *Club) AddEvent(event *Event) {
	c.events[event.id] = event
}

func (c *Club) RemoveEvent(event *Event) {
	delete(c.events, event.id)
}
