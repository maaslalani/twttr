package views

// View is a type that represents a view in the application.
// See below for the different views a user can navigate to.
type View int

const (
	// Loading is the view when the model is fetching data from the
	// server, in this case the Twitter API.
	Loading View = iota
	// Help is the view to show when the user requests to view all of the
	// keybindings.
	Help
	// Compose is the view to show when the user is composing a new tweet.
	Compose
	// Home is the view to show when the user is viewing their home
	// timeline.
	Home
)
