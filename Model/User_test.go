package Model

/*
	TestGetRsvpList	Returns the user's RSVP list
	TestRsvpAdd	Adds an event ID to the user's RSVP list
	TestRsvpAddUpdatesEvent	Adds the user's ID to the event's attendees
	TestRsvpRemove	Removes an event ID from the user's RSVP list
	TestRsvpRemoveUpdatesEvent	Removes the user's ID from the event's attendees
	TestRsvpSameEventTwice	RSVPing to the same event doesn't create duplicate entries
	TestRsvpRemoveNonRsvpEvent	Removing an event the user hasn't RSVP'd to doesn't cause problems
	TestCreateEvent	Creates an event with the correct properties
	TestCreateEventGeneratesID	Created event receives the next Campus event ID
	TestCreateEventAddedToClub	Created event is added to the specified club
	TestCreateMultipleEvents	Multiple events receive unique IDs
	TestDeleteEvent	Deletes an event from its club
*/