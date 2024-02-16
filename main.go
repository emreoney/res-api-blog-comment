package main

import (
	"gomod/database"
	"gomod/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	database.Init()

	router := mux.NewRouter()

	router.HandleFunc("/", handlers.HandlerHomePage)

	// Blog(POSTS) APIs
	router.HandleFunc("/posts", handlers.HandlerCreatePost).Methods("POST")
	router.HandleFunc("/posts", handlers.HandlerGetPosts).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.HandlerGetPost).Methods("GET")
	router.HandleFunc("/posts/{id}", handlers.HandlerUpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", handlers.HandlerDeletePost).Methods("DELETE")

	// Comment APIs
	router.HandleFunc("/posts/{id}/comments", handlers.HandlerCreateComment).Methods("POST")
	router.HandleFunc("/posts/{id}/comments", handlers.HandlerGetCommentsForSpesificBlog).Methods("GET")
	router.HandleFunc("/posts/{id}/comments/{commentId}", handlers.HandlerGetCommentForSpesificBlog).Methods("GET")
	router.HandleFunc("/posts/{id}/comments/{commentId}", handlers.HandlersUpdateComment).Methods("PUT")
	router.HandleFunc("/posts/{id}/comments/{commentId}", handlers.HandlerDeleteComment).Methods("DELETE")

	http.ListenAndServe(":8080", router)

}
