// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Assistant struct {
	ID      string `json:"id"`
	Initial string `json:"initial"`
	Name    string `json:"name"`
}

type FindAssistantByID struct {
	ID string `json:"id"`
}

type FindAssistantByInitial struct {
	Initial string `json:"initial"`
}

type FindTodoByID struct {
	ID string `json:"id"`
}

type MarkTodoAsDone struct {
	ID   string `json:"id"`
	Done bool   `json:"done"`
}

type NewAssistant struct {
	Initial string `json:"initial"`
	Name    string `json:"name"`
}

type NewTodo struct {
	Text             string `json:"text"`
	AssistantInitial string `json:"assistantInitial"`
}
