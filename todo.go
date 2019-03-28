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

// String returns stirng representation of status.
func (s Status) String() string {
	switch s {
	case Pending:
		return "Pending"
	case Done:
		return "Done"
	case InProgress:
		return "InProgress"
	default:
		return ""
	}
}

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
	// All returns all todos.
	All() []*Todo
	// ByID retruns todo by its id.
	ByID(id string) (*Todo, error)
	// Create creates new todo.
	Create(el *Todo) error
	// Status changes todo item status.
	Status(todoID string, status Status) error
}
