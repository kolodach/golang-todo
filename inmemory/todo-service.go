package inmemory

import todo "github.com/kolodach/golang-todo"

// Ensure TodoService implements todo.TodoService.
var _ todo.TodoService = &TodoService{}

// TodoService represents a service for managing todo
type TodoService struct {
	Stor *Store
}

// NewTodoService creates new instance of TodoService and
// returns pointer on it.
func NewTodoService() *TodoService {
	return &TodoService{
		Stor: New(),
	}
}

// TodoById gets specifict Todo by its id.
func (ts TodoService) TodoByID(id string) (*todo.Todo, error) {
	el, err := ts.Stor.Get(id)
	if err != nil {
		return nil, err
	}
	todoItem, _ := el.(*todo.Todo)
	return todoItem, nil
}

// Create inserts given element to store
func (ts TodoService) Create(el *todo.Todo) error {
	item := interface{}(el)
	err := ts.Stor.Add(el.ID, item)
	if err != nil {
		return err
	}
	return nil
}

// SetStatus changes current todo status to new.
func (ts TodoService) SetStatus(todoID string, status todo.Status) error {
	el, err := ts.Stor.Get(todoID)
	if err != nil {
		return err
	}
	todoEl, _ := el.(*todo.Todo)
	todoEl.Status = status
	return nil
}
