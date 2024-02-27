package controllers

import (
	"encoding/json"
	"http-server/src/models"
	"http-server/src/structs"
	"http-server/src/utils"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func CreatePost(db *gorm.DB, post models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		createPostInput := structs.CreatePost{}

		// check if user is logged in
		if !utils.VerifyAuthState(w, r) {
			utils.SetResponse(w, http.StatusForbidden, "error", "You need to login")
			return
		}

		// validate request body
		if err := json.NewDecoder(r.Body).Decode(&createPostInput); err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		// trimming input data for blank spaces
		title := strings.Trim(createPostInput.Title, " ")
		body := strings.Trim(createPostInput.Body, " ")

		// basic input validation
		if len(title) < 3 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Title must be at least 3 characters long", "field": "title"})
			return
		}
		if len(body) < 5 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Body must be at least 5 characters long", "field": "body"})
			return
		}

		cookie, _ := r.Cookie("JWT")
		token, _ := utils.VerifyJWT(cookie.Value)

		newPost := models.Post{
			ID:       uuid.New(),
			Title:    title,
			Body:     body,
			AuthorID: token.Claims.(*structs.UserClaims).UserID,
		}

		db.Create(&newPost)

		var author models.User

		db.First(&author, "id = ?", newPost.AuthorID)

		utils.SetResponse(w, http.StatusCreated, "success", map[string]any{
			"msg": "Post successfully created",
			"user": map[string]any{
				"id":         newPost.ID,
				"title":      newPost.Title,
				"body":       newPost.Body,
				"createdAt":  newPost.CreatedAt,
				"authorId":   newPost.AuthorID,
				"authorName": author.Name,
			},
		})
	}
}

func UpdatePostById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var postToBeUpdated models.Post
		updatePostInput := structs.UpdatePost{}

		// check if user is logged in
		if !utils.VerifyAuthState(w, r) {
			utils.SetResponse(w, http.StatusUnauthorized, "error", "You need to login")
			return
		}

		// validate request body
		if err := json.NewDecoder(r.Body).Decode(&updatePostInput); err != nil {
			utils.SetResponse(w, http.StatusBadRequest, "error", err.Error())
			return
		}

		// getting post id from url
		vars := mux.Vars(r)
		result := db.Find(&postToBeUpdated, "id = ?", vars["postID"])

		// checking if post exists or not
		if result.Error != nil {
			utils.SetResponse(w, http.StatusNotFound, "error", "Requested post not found")
			return
		}

		// checking if post is deleted or not
		if !postToBeUpdated.DeletedAt.IsZero() {
			utils.SetResponse(w, http.StatusForbidden, "error", "This post has been deleted")
			return
		}

		// trimming input data for blank spaces
		newTitle := strings.Trim(updatePostInput.Title, " ")
		newBody := strings.Trim(updatePostInput.Body, " ")

		// basic input validation
		if len(newTitle) < 3 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Title must be at least 3 characters long", "field": "newTitle"})
			return
		}
		if len(newBody) < 5 {
			utils.SetResponse(w, http.StatusBadRequest, "error", map[string]string{"value": "Body must be at least 5 characters long", "field": "body"})
			return
		}

		// check if post.authorId matches logged in userID
		cookie, _ := r.Cookie("JWT")
		token, _ := utils.VerifyJWT(cookie.Value)

		if postToBeUpdated.AuthorID != token.Claims.(*structs.UserClaims).UserID {
			utils.SetResponse(w, http.StatusForbidden, "error", "You are not authorized to update this post")
			return
		}

		db.Model(&models.Post{}).Where("id = ?", postToBeUpdated.ID).Updates(models.Post{Title: newTitle, Body: newBody})

		var author models.User

		db.Select("name").Take(&author, "id = ?", postToBeUpdated.AuthorID)

		utils.SetResponse(w, http.StatusCreated, "success", map[string]any{
			"msg": "Post successfully updated",
			"user": map[string]any{
				"id":         postToBeUpdated.ID,
				"title":      postToBeUpdated.Title,
				"body":       postToBeUpdated.Body,
				"createdAt":  postToBeUpdated.CreatedAt,
				"authorId":   postToBeUpdated.AuthorID,
				"authorName": author.Name,
			},
		})
	}
}

func GetAllPosts(db *gorm.DB, posts []models.Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var allPosts []structs.Post

		db.Select("posts.*, users.name AS author_name").Joins("LEFT JOIN users ON users.id=posts.author_id").Order("created_at DESC").Where("posts.deleted_at IS NULL").Find(&allPosts)

		utils.SetResponse(w, http.StatusOK, "success", allPosts)
	}
}

func GetPostById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// getting post id from url
		vars := mux.Vars(r)

		var post models.Post
		result := db.Select("posts.*, users.name AS author_name").Where("posts.id = ?", vars["postID"]).Joins("LEFT JOIN users ON users.id=posts.author_id").Take(&post)

		// checking if post exists or not
		if result.Error != nil {
			utils.SetResponse(w, http.StatusNotFound, "error", "Requested post not found")
			return
		}

		// checking if post is deleted or not
		if !post.DeletedAt.IsZero() {
			utils.SetResponse(w, http.StatusForbidden, "error", "This post has been deleted")
			return
		}

		var resPost structs.Post

		resPost.ID = post.ID
		resPost.Title = post.Title
		resPost.Body = post.Body
		resPost.CreatedAt = post.CreatedAt
		resPost.UpdatedAt = post.UpdatedAt
		resPost.AuthorID = post.AuthorID
		resPost.AuthorName = *post.AuthorName

		utils.SetResponse(w, http.StatusOK, "success", resPost)
	}
}

func DeletePostById(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var postToBeDeleted models.Post

		// check if user is logged in
		if !utils.VerifyAuthState(w, r) {
			utils.SetResponse(w, http.StatusUnauthorized, "error", "You need to login")
			return
		}

		// getting post id from url
		vars := mux.Vars(r)

		result := db.Find(&postToBeDeleted, "id = ?", vars["postID"])

		// checking if post exists or not
		if result.Error != nil {
			utils.SetResponse(w, http.StatusNotFound, "error", "Requested post not found")
			return
		}

		// check if post is already deleted or not
		if len(postToBeDeleted.DeletedAt.String()) > 6 {
			utils.SetResponse(w, http.StatusForbidden, "error", "This post has already been deleted")
			return
		}

		/// check if post.authorId matches logged in userID
		cookie, _ := r.Cookie("JWT")
		token, _ := utils.VerifyJWT(cookie.Value)

		if postToBeDeleted.AuthorID != token.Claims.(*structs.UserClaims).UserID {
			utils.SetResponse(w, http.StatusForbidden, "error", "You are not authorized to delete this post")
			return
		}

		db.Where("posts.id = ?", vars["postID"]).Delete(&postToBeDeleted)

		utils.SetResponse(w, http.StatusOK, "success", "Post successfully deleted")
	}
}
