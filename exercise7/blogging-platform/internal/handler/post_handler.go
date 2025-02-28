package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/service"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{
		service: service,
	}
}

// CreatePost 
func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post models.Post

	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	if err := h.service.CreatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Error when sending the response", http.StatusInternalServerError)
	}
}

// GetAllPosts 
func (h *PostHandler) GetAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		http.Error(w, "Error receiving posts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, "Error when sending the response", http.StatusInternalServerError)
	}
}

// GetPostByID 
func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	post, err := h.service.GetPostByID(int64(id))
	if err != nil {
		http.Error(w, "The post was not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(post); err != nil {
		http.Error(w, "Error when sending the response", http.StatusInternalServerError)
	}
}

// UpdatePost 
func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	var post models.Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	post.ID = int64(id)

	if err := h.service.UpdatePost(&post); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The post has been updated successfully"))
}

// DeletePost — удаление поста по ID
func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	if err := h.service.DeletePost(int64(id)); err != nil {
		http.Error(w, "Error when deleting a post", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("deleted successfully"))
}
