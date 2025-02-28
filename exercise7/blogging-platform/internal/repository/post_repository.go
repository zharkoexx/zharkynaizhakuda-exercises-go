package repository

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"github.com/talgat-ruby/exercises-go/exercise7/blogging-platform/models"
)

type PostRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{db: db}
}

// Create 
func (r *PostRepository) Create(post *models.Post) error {
	tagsString := strings.Join(post.Tags, ",")
	query := `INSERT INTO posts (title, content, category, tags) VALUES (?, ?, ?, ?)`
	_, err := r.db.Exec(query, post.Title, post.Content, post.Category, tagsString)
	if err != nil {
		log.Printf("Error when creating a post: %v", err)
	}
	return err
}

// GetPostByID 
func (r *PostRepository) GetPostByID(id int) (*models.Post, error) {
	query := `SELECT id, title, content, created_at, category, tags, updated_at FROM posts WHERE id = ?`
	row := r.db.QueryRow(query, id)

	post := &models.Post{}
	var tagsString sql.NullString
	var createdAtStr, updatedAtStr string
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&createdAtStr,
		&post.Category,
		&tagsString,
		&updatedAtStr,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Post with ID %d  not found", id)
			return nil, nil
		}
		log.Printf("Error when receiving a post by ID: %v", err)
		return nil, err
	}

	post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
	if err != nil {
		log.Printf("Error during conversion created_at: %v", err)
		return nil, err
	}

	post.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
	if err != nil {
		log.Printf("Error during conversion updated_at: %v", err)
		return nil, err
	}

	if tagsString.Valid {
		post.Tags = strings.Split(tagsString.String, ",")
	}

	return post, nil
}

// GetAll 
func (r *PostRepository) GetAll() ([]models.Post, error) {
	query := `SELECT id, title, content, created_at, category, tags, updated_at FROM posts`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("Error when receiving all posts: %v", err)
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		post := models.Post{}
		var tagsString sql.NullString
		var createdAtStr, updatedAtStr string
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&createdAtStr,
			&post.Category,
			&tagsString,
			&updatedAtStr,
		)
		if err != nil {
			log.Printf("Error when scanning a line of a post: %v", err)
			return nil, err
		}

		post.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			log.Printf("Error during conversion created_at: %v", err)
			return nil, err
		}

		post.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
		if err != nil {
			log.Printf("Error during conversion updated_at: %v", err)
			return nil, err
		}

		if tagsString.Valid {
			post.Tags = strings.Split(tagsString.String, ",")
		}

		posts = append(posts, post)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error when processing strings: %v", err)
		return nil, err
	}

	return posts, nil
}

// Update 
func (r *PostRepository) Update(post *models.Post) error {
	tagsString := strings.Join(post.Tags, ",")
	query := `UPDATE posts SET title = ?, content = ?, category = ?, tags = ?, updated_at = CURRENT_TIMESTAMP WHERE id = ?`
	_, err := r.db.Exec(query, post.Title, post.Content, post.Category, tagsString, post.ID)
	if err != nil {
		log.Printf("Error when updating a post with an ID %d: %v", post.ID, err)
	}
	return err
}

// Delete 
func (r *PostRepository) Delete(id int) error {
	query := `DELETE FROM posts WHERE id = ?`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("Error when deleting a post with an ID %d: %v", id, err)
	}
	return err
}
