package handler

import (
	"encoding/json"
	"fmt"
	"log"
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

func (h *Handler) Refresh() {
	h.Posts = make(chan models.Post)
	h.Comments = make(chan models.Comment)
	h.Users = make(chan models.User)
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandleRoute(w http.ResponseWriter, r *http.Request) {
	h.Refresh()
	db := database.NewDatabase()
	errorChan := make(chan error)
	wg := sync.WaitGroup{}
	go func() {
		for i := range errorChan {
			log.Fatal("an error occured during fetching:", i)
		}
	}()
	wg.Add(3)
	go func() {
		defer wg.Done()
		for v := range h.Users {
			db.Users[int(v.Id)] = v
		}
	}()

	go func() {
		defer wg.Done()
		for v := range h.Comments {
			db.Comments[int(v.PostID)]++
		}
	}()
	go func() {
		defer wg.Done()
		for v := range h.Posts {
			db.Posts[int(v.ID)] = v
		}
	}()

	wg.Add(3)
	go func() {
		defer wg.Done()
		defer close(h.Posts)
		helper.FetchPosts(h.Posts, errorChan)
	}()

	go func() {
		defer wg.Done()
		defer close(h.Comments)
		helper.FetchComments(h.Comments, errorChan)
	}()

	go func() {
		defer wg.Done()
		defer close(h.Users)
		helper.FetchUsers(h.Users, errorChan)
	}()

	wg.Wait()
	close(errorChan)

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
