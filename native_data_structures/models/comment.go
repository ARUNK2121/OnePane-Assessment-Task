package models

type Comment struct {
	PostID int64  `json:"postId"`
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Body   string `json:"body"`
}
