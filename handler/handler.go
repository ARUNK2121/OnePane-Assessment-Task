package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	database "one_pane_assessment/Database"
	"one_pane_assessment/helper"
	"one_pane_assessment/models"
)

type Handler struct {
	Posts    []models.Post
	Comments []models.Comment
	Users    []models.User
}

func NewHandler() *Handler {
	return &Handler{
		Posts:    make([]models.Post, 0),
		Comments: make([]models.Comment, 0),
		Users:    make([]models.User, 0),
	}
}

func (h *Handler) HandleRoute(w http.ResponseWriter, r *http.Request) {
	db := database.NewDatabase()
	var err error
	h.Posts, err = helper.FetchPosts()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching posts: %v", err), http.StatusInternalServerError)
		return
	}
	h.Comments, err = helper.FetchComments()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching comments: %v", err), http.StatusInternalServerError)
		return
	}
	h.Users, err = helper.FetchUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
		return
	}

	// Store to map for easy manipulation and better time complexity for fetching (in worst case O(n))
	for _, v := range h.Users {
		db.Users[int(v.Id)] = v
	}

	for _, v := range h.Comments {
		db.Comments[int(v.PostID)]++
	}

	for _, v := range h.Posts {
		db.Posts[int(v.ID)] = v
	}

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
