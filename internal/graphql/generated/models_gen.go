// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

type NewTodo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"isCompleted"`
}
