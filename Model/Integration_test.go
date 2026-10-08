package Model

/*
	TestCreateEventAndRSVP	User creates an event and successfully RSVPs
	TestRSVPBidirectionalConsistency	RSVP appears in both the User and Event
	TestRemoveRSVPBidirectionalConsistency	Removing RSVP removes it from both sides
	TestCreateEventAppearsOnCampus	A created event can be found through Campus.getEvents()
	TestDeleteEventRemovesFromCampus	Deleting an event makes it disappear from Campus's event collection
	TestMultipleUsersRSVP	Multiple users can RSVP to the same event
	TestMultipleUsersAndEvents	Users can RSVP to different combinations of events without affecting each other
*/