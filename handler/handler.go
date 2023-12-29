package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "one_pane_assessment/database"
	"one_pane_assessment/helper"
	"one_pane_assessment/models"
	"sync"
)

type Handler struct {
	Posts    chan models.Post
	Comments chan models.Comment
	Users    chan models.User
}

func NewHandler() *Handler {
	return &Handler{
		Posts:    make(chan models.Post),
		Comments: make(chan models.Comment),
		Users:    make(chan models.User),
	}
}

func (h *Handler) HandleRoute(w http.ResponseWriter, r *http.Request) {
	db := database.NewDatabase()
	var errorChan chan error

	wg := sync.WaitGroup{}
	wg.Add(3)
	go func(c <-chan models.User, wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range h.Users {
			db.Users[int(v.Id)] = v
		}
	}(h.Users, &wg)

	go func(c <-chan models.Comment, wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range h.Comments {
			db.Comments[int(v.PostID)]++
		}
	}(h.Comments, &wg)

	go func(c <-chan models.Post, wg *sync.WaitGroup) {
		defer wg.Done()
		for v := range h.Posts {
			db.Posts[int(v.ID)] = v
		}
	}(h.Posts, &wg)

	go helper.FetchPosts(h.Posts, errorChan)
	go helper.FetchComments(h.Comments, errorChan)
	go helper.FetchUsers(h.Users, errorChan)

	wg.Wait()

	var Result []models.Result
	for i, v := range db.Posts {
		Result = append(Result, models.Result{
			PostID:        i,
			PostName:      v.Title,
			CommentsCount: db.Comments[i],
			UserName:      db.Users[int(v.UserID)].Name,
			Body:          v.Body,
		})
	}

	jsonResponse, err := json.Marshal(Result)
	if err != nil {
		http.Error(w, fmt.Sprintf("error marshaling json: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
