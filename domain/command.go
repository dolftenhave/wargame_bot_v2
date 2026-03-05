package domain

// Command represends a command that war recieved to interact with the wargame server.
type Command struct {
	SenderID int
	Source string
	Content string
}
