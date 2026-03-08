package command

// Rends the command result back to the sender.
type Responder interface {
	SendResult(result CommandResult) error
}
