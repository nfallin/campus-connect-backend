package Model

// default overload: gets every event
func viewEvents(c Campus) map[int]*Event {
	return c.getEvents()
}