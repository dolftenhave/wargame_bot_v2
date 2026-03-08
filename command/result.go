package command

// The type of the result
// - ResultMessage
// - ResultList
// - ResultError
// - ResultSuccess
type ResultType int

const (
	// A simple text message result
	ResultMessage ResultType = iota
	// A list of items (maps, modes, players, commands ect.)
	ResultList
	// An error
	ResultError
	// The action was successful
	ResultSuccess
)

// Source agnostic result data
type CommandResult struct {
	Type ResultType
	Title string
	Message string
	Fields []ResultField
	Error error
}

// A key value set of result data.
type ResultField struct {
	Name string
	Value string
}
