package entities

// TodoList ...
type TodoList struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// TodoItem ...
type TodoItem struct {
	ID          int    `json:"-"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}

// UsersList ...
type UsersList struct {
	ID     int
	UserID int
	ListID int
}

// ListItems ...
type ListItems struct {
	ID     int
	ListID int
	ItemID int
}
