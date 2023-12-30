package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"one_pane_assessment/models"
)

func FetchPosts(c chan models.Post, errorChan chan error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		errorChan <- err
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("failed to fetch posts. status code: %d", res.StatusCode)
		return
	}

	var posts []models.Post
	err = json.NewDecoder(res.Body).Decode(&posts)
	if err != nil {
		errorChan <- err
		return
	}

	for _, v := range posts {
		v := v
		c <- v
	}

}

func FetchComments(c chan models.Comment, errorChan chan error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/comments")
	if err != nil {
		errorChan <- err
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("failed to fetch comments. status code: %d", res.StatusCode)
		return
	}

	var comments []models.Comment
	err = json.NewDecoder(res.Body).Decode(&comments)
	if err != nil {
		errorChan <- err
		return
	}

	for _, v := range comments {
		v := v
		c <- v
	}
}

func FetchUsers(c chan models.User, errorChan chan error) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		errorChan <- err
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		errorChan <- fmt.Errorf("failed to fetch users. status code: %d", res.StatusCode)
		return
	}

	var users []models.User
	err = json.NewDecoder(res.Body).Decode(&users)
	if err != nil {
		errorChan <- err
		return
	}

	for _, v := range users {
		v := v
		c <- v
	}
}
