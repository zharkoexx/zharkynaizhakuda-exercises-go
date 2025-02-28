package service

import (
	"errors"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/internal/repository"
	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostService struct {
	repo *repository.PostRepository
}

func NewPostService(repo *repository.PostRepository) *PostService {
	return &PostService{repo: repo}
}

// CreatePost — проверка данных и создание поста
func (s *PostService) CreatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" || post.Category == "" {
		return errors.New("all fields (title, content, category) are required")
	}
	return s.repo.Create(post)
}

// GetAllPosts — получение всех постов
func (s *PostService) GetAllPosts() ([]models.Post, error) {
	return s.repo.GetAll()
}

// GetPostByID — получение поста по ID
func (s *PostService) GetPostByID(id int64) (*models.Post, error) {
	post, err := s.repo.GetPostByID(int(id))
	if err != nil {
		return nil, errors.New("the post was not found")
	}
	return post, nil
}

// UpdatePost — обновление поста по ID
func (s *PostService) UpdatePost(post *models.Post) error {
	if post.Title == "" || post.Content == "" || post.Category == "" {
		return errors.New("all fields (title, content, category) are required")
	}

	existingPost, err := s.repo.GetPostByID(int(post.ID))
	if err != nil {
		return errors.New("the post was not found")
	}

	existingPost.Title = post.Title
	existingPost.Content = post.Content
	existingPost.Category = post.Category
	existingPost.Tags = post.Tags

	return s.repo.Update(existingPost)
}

// DeletePost — удаление поста по ID
func (s *PostService) DeletePost(id int64) error {
	_, err := s.repo.GetPostByID(int(id))
	if err != nil {
		return errors.New("the post was not found")
	}
	return s.repo.Delete(int(id))
}
