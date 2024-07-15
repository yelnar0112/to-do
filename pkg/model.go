package todo

type TodoList struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	ActiveAt  string `json:"activeAt"`
	Completed bool   `json:"completed"`
}
