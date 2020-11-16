// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewTodo struct {
	Text     string `json:"text"`
	Username string `json:"username"`
}

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	Username string  `json:"username"`
	Todos    []*Todo `json:"todos"`
}

func (User) IsEntity() {}
