package todo

// Todo statuses
const (
	Pending    = Status(iota)
	Done       = Status(iota)
	InProgress = Status(iota)
)

// Common errors.
const (
	ErrItemExists = Error("Item with given ID already exists")
	ErrTodoExists = Error("Todo already exists")
)

// Status represents current state for specifict task.
type Status int8

// Error represents todo error.
type Error string

// Status return the status integer id value.
func (s Status) Status() int8 { return int8(s) }

// Error returns the string error message.
func (e Error) Error() string { return string(e) }

// Todo represents specific task in agenda.
// It can be in certain state which represents task progress.
type Todo struct {
	ID     string
	Name   string
	Status Status
}

// TodoService manages todo item.
type TodoService interface {
	TodoByID(id string) (*Todo, error)
	Create(el *Todo) error
	SetStatus(todoID string, status Status) error
}
