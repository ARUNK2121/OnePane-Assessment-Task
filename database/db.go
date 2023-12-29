package database

import "one_pane_assessment/models"

type Database struct {
	Posts    map[int]models.Post
	Comments map[int]int
	Users    map[int]models.User
}

func NewDatabase() *Database {
	return &Database{
		Posts:    make(map[int]models.Post),
		Comments: map[int]int{},
		Users:    make(map[int]models.User),
	}
}
