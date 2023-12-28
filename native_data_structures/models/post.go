package models

type Post struct {
	UserID int64  `json:"userId"`
	ID     int64  `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

type Result struct {
	PostID        int
	PostName      string
	CommentsCount int
	UserName      string
	Body          string
}
