package todo

type TodoList struct {
	Id          int `json:"id"`
	Title       int `json:"title"`
	Description int `json:"description"`
}

type UsersLust struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id          int  `json:"id"`
	Title       int  `json:"title"`
	Description int  `json:"description"`
	Done        bool `json:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
