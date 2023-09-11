package model

import "time"

type Todo struct {
	Id        int       `json:"todo"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo
