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

func HandlerCreateComment(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	var newComment models.Comment
	newComment.BlogID = blogId
	json.NewDecoder(r.Body).Decode(&newComment)

	database.DB.Create(&newComment)

	data, err := json.Marshal(newComment)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))
}
func HandlerGetCommentsForSpesificBlog(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)

	var comments []models.Comment
	database.DB.Find(&comments, "blog_id = ?", blogId)

	data, err := json.Marshal(comments)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))
}
func HandlerGetCommentForSpesificBlog(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)
	commentId, err := strconv.Atoi(variables["commentId"])

	var comment models.Comment
	database.DB.Find(&comment, "blog_id = ? AND id = ?", blogId, commentId)

	data, err := json.Marshal(comment)
	helpers.CheckError(err)
	fmt.Fprintf(w, string(data))
}
func HandlersUpdateComment(w http.ResponseWriter, r *http.Request) {
	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)
	commentId, err := strconv.Atoi(variables["commentId"])

	var updatedComment models.Comment
	updatedComment.ID = uint(commentId)
	updatedComment.BlogID = blogId
	json.NewDecoder(r.Body).Decode(&updatedComment)

	database.DB.Save(&updatedComment)

	data, err := json.Marshal(updatedComment)
	fmt.Fprintf(w, string(data))

}
func HandlerDeleteComment(w http.ResponseWriter, r *http.Request) {
	var deletedComment models.Comment

	variables := mux.Vars(r)
	blogId, err := strconv.Atoi(variables["id"])
	helpers.CheckError(err)
	commentId, err := strconv.Atoi(variables["commentId"])

	deletedComment.ID = uint(commentId)
	deletedComment.BlogID = blogId

	database.DB.Delete(&deletedComment)

	responseMessage := models.Message{"Comment has deleted"}
	data, err := json.Marshal(responseMessage)
	fmt.Fprintf(w, string(data))
}
