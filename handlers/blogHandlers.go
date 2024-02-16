package handlers

import (
	"encoding/json"
	"fmt"
	"gomod/database"
	"gomod/helpers"
	"gomod/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func HandlerCreatePost(w http.ResponseWriter, r *http.Request) {
	var newBlog models.Blog
	json.NewDecoder(r.Body).Decode(&newBlog)

	database.DB.Create(&newBlog)

	data, err := json.Marshal(newBlog)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))
}

func HandlerGetPosts(w http.ResponseWriter, r *http.Request) {
	var blogs []models.Blog
	//database.DB.Find(&blogs)
	database.DB.Preload("Comments").Find(&blogs)

	data, err := json.Marshal(blogs)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))
}

func HandlerGetPost(w http.ResponseWriter, r *http.Request) {
	var blog models.Blog
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	blog.ID = uint(blogId)
	database.DB.Preload("Comments").First(&blog)

	data, err := json.Marshal(blog)
	fmt.Fprintf(w, string(data))
}

func HandlerUpdatePost(w http.ResponseWriter, r *http.Request) {
	var updatedBlog models.Blog
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	updatedBlog.ID = uint(blogId)
	json.NewDecoder(r.Body).Decode(&updatedBlog)
	database.DB.Save(&updatedBlog)

	data, err := json.Marshal(updatedBlog)
	fmt.Fprintf(w, string(data))

}
func HandlerDeletePost(w http.ResponseWriter, r *http.Request) {
	var deletedBlog models.Blog
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	deletedBlog.ID = uint(blogId)
	database.DB.Delete(&deletedBlog)

	responseMessage := models.Message{"Data has deleted"}
	data, err := json.Marshal(responseMessage)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))

}
