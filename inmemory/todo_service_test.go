package inmemory_test

import (
	"errors"
	"testing"

	todo "github.com/kolodach/golang-todo"
	"github.com/kolodach/golang-todo/inmemory"
)

// Ensure can add todo.
func TestTodoService_CreateTodo(t *testing.T) {
	s := inmemory.TodoService{
		Stor: inmemory.New(),
	}
	el := &todo.Todo{
		ID:     "1",
		Name:   "test",
		Status: todo.Status(todo.Pending),
	}

	s.Create(el)

	savedEl, err := s.ByID("1")
	if err != nil {
		t.Fatal(err)
	}
	if savedEl == nil {
		t.Fatal(errors.New("el not exists"))
	}
	if *savedEl != *el {
		t.Fatal(errors.New("Pointers are different"))
	}
	if savedEl.ID != "1" || savedEl.Name != "test" || savedEl.Status != todo.Status(todo.Pending) {
		t.Fatal(errors.New("invalid todo data"))
	}
}

// Ensure TodoService can change todo status.
func TestTodoService_ChangeTodoState(t *testing.T) {
	s := inmemory.TodoService{
		Stor: inmemory.New(),
	}
	el := &todo.Todo{
		ID:     "1",
		Name:   "test",
		Status: todo.Status(todo.Pending),
	}

	s.Create(el)
	statErr := s.Status(el.ID, todo.Status(todo.Done))

	if statErr != nil {
		t.Fatal(statErr)
	}
	savedEl, err := s.ByID("1")
	if err != nil {
		t.Fatal(err)
	}
	if savedEl == nil {
		t.Fatal(errors.New("el not exists"))
	}
	if *savedEl != *el {
		t.Fatal(errors.New("Pointers are different"))
	}
	if savedEl.Status != todo.Status(todo.Done) {
		t.Fatal(errors.New("Todo status is not valid"))
	}
}

// Ensure TodoService can return all elements
func TestTodoService_CanGetAll(t *testing.T) {
	s := inmemory.NewTodoService()

	el1 := &todo.Todo{
		ID:     "1",
		Name:   "test1",
		Status: todo.Status(todo.Pending),
	}
	el2 := &todo.Todo{
		ID:     "2",
		Name:   "test2",
		Status: todo.Status(todo.Pending),
	}

	s.Create(el1)
	s.Create(el2)

	all := s.All()
	if all == nil {
		t.Fatal(errors.New("null slice"))
	}
	if len(all) != 2 {
		t.Fatal(errors.New("Invalid lingth"))
	}
	if all[0] != el1 {
		t.Fatal(errors.New("Invalid element1"))
	}
	if all[1] != el2 {
		t.Fatal(errors.New("Invalid element2"))
	}
}
