package domain

type Todo struct {
	Title TodoTitle `json:"title"`
	Done TodoDone `json:"done"`
}

type TodoId struct {
	Value int `json:"value"`
}

type TodoTitle struct {
	Value string `json:"value"`
}

type TodoDone struct {
	Value bool `json:"value"`
}

func NewTodoTitle(value string) TodoTitle {
	return TodoTitle{value}
}

func NewTodoDone(value bool) TodoDone {
	return TodoDone{value}
}

func NewTodo(title string, done bool) Todo {
	return Todo{
		Title: NewTodoTitle(title),
		Done: NewTodoDone(done),
	}
}