package Model

/*
	TestNewCampus	Campus initializes with empty user/club maps and all ID counters starting at 1
	TestGetNextUserID	User IDs increment correctly
	TestGetNextClubID	Club IDs increment correctly
	TestGetNextEventID	Event IDs increment correctly
	TestCreateUser	Creates a user with the correct name, ID, and initialized RSVP list
	TestCreateMultipleUsers	Multiple users receive unique sequential IDs
	TestCreateClub	Creates a club with the correct name, ID, and initialized event map
	TestCreateMultipleClubs	Multiple clubs receive unique sequential IDs
	TestAddUser	Adds a user to the Campus user map
	TestRemoveUser	Removes a user from the Campus user map
	TestAddClub	Adds a club to the Campus club map
	TestRemoveClub	Removes a club from the Campus club map
	TestGetUsers	Returns the Campus users
	TestGetClubs	Returns the Campus clubs
	TestGetEventsEmpty	Returns an empty map when there are no events
	TestGetEvents	Returns all events across all clubs
	TestGetEventsMultipleClubs	Correctly combines events from multiple clubs
	TestGetEventsNoDuplicateIDs	Events are correctly represented by their IDs
*/