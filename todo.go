package todo

import "errors"

type TodoList struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Done  bool   `json:"done" db:"done"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
	Done  bool   `json:"done" db:"done"`
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}

type UpdateListInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

type UpdateItemInput struct {
	Title *string `json:"title"`
	Done  *bool   `json:"done"`
}

func (i UpdateItemInput) Validate() error {
	if i.Title == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

func (i UpdateListInput) Validate() error {
	if i.Title == nil && i.Done == nil {
		return errors.New("update structure has no values")
	}
	return nil
}
