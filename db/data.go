package db

import (
    "fmt"
    "github.com/murilo-toddy/mess/model/todo"
)

var currentId int

var todos todo.Todos

func init() {
	DataTodoCreate(todo.Todo{Name: "Task 1"})
	DataTodoCreate(todo.Todo{Name: "Task 2"})
}

func DataTodoFind(id int) todo.Todo {
	for _, todo := range todos {
		if todo.Id == id {
			return todo
		}
	}
	// not found
	return todo.Todo{}
}

func DataTodoCreate(todo todo.Todo) todo.Todo {
	currentId += 1
	todo.Id = currentId
	todos = append(todos, todo)
	return todo
}

func DataTodoRemove(id int) error {
	for i, todo := range todos {
		if todo.Id == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find task with id %d", id)
}
