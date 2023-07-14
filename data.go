package main

import "fmt"

var currentId int

var todos Todos

func init() {
    DataTodoCreate(Todo{Name: "Task 1"})
    DataTodoCreate(Todo{Name: "Task 2"})
}

func DataTodoFind(id int) Todo {
    for _, todo := range todos {
        if todo.Id == id {
            return todo
        }
    }
    // not found
    return Todo{}
}

func DataTodoCreate(todo Todo) Todo {
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

