package common

import (
	"errors"
	"fmt"
)

type Todo struct {
	ID       int    `json:"id"`
	Desc     string `json:"desc"`
	Status   string `json:"status"`
	Modified string `json:"modified"`
}

type TodoList []Todo

func (t TodoList) Len() int {
	return len(t)
}

func (t TodoList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TodoList) Less(i, j int) bool {
	return t[i].ID < t[j].ID
}

func (t TodoList) FindByID(id int) (int, error) {
	for i, todo := range t {
		if todo.ID == id {
			return i, nil
		}
	}
	errmsg := fmt.Sprintf("todo: Cannot find a todo item with id \"%d\"", id)
	return 0, errors.New(errmsg)
}

func (t TodoList) Remove(fn func(td Todo) bool) TodoList {
	var result TodoList
	for _, todo := range t {
		if !fn(todo) {
			result = append(result, todo)
		}
	}
	return result
}
