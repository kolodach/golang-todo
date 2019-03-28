package inmemory

import (
	"sort"

	todo "github.com/kolodach/golang-todo"
)

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

// All returns all todos.
func (ts TodoService) All() []*todo.Todo {
	els := ts.Stor.All()
	if els == nil {
		return nil
	}
	todos := make([]*todo.Todo, len(els))
	for i, el := range els {
		todos[i] = el.(*todo.Todo)
	}
	sort.Slice(todos, func(i, j int) bool {
		return todos[i].Name > todos[j].Name
	})
	return todos
}

// TodoById gets specifict Todo by its id.
func (ts TodoService) ByID(id string) (*todo.Todo, error) {
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
func (ts TodoService) Status(ID string, s todo.Status) error {
	el, err := ts.Stor.Get(ID)
	if err != nil {
		return err
	}
	todoEl, _ := el.(*todo.Todo)
	todoEl.Status = s
	return nil
}
