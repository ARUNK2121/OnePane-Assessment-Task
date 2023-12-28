package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"one_pane_assessment/models"
)

func FetchPosts() ([]models.Post, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch posts. status code: %d", res.StatusCode)
	}

	var posts []models.Post
	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func FetchComments() ([]models.Comment, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch comments. status code: %d", res.StatusCode)
	}

	var comments []models.Comment
	err = json.NewDecoder(res.Body).Decode(&comments)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func FetchUsers() ([]models.User, error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch users. status code: %d", res.StatusCode)
	}

	var users []models.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
