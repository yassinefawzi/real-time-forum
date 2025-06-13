package handlers

import (
	"database/sql"
	"encoding/json"
	g "forum/server/global"
	"log"
	http "net/http"
	"github.com/google/uuid"

	_ "github.com/mattn/go-sqlite3"
)

func insertPost(db *sql.DB, post g.Post) error {
	query := `
		INSERT INTO posts (id, title, content, category)
		VALUES (?, ?, ?, ?);
	`

	_, err := db.Exec(query, post.ID, post.Title, post.Content, post.Category)

	return err
}



func Getcreatepost(w http.ResponseWriter, r *http.Request) {
	// Parse the form data
	if err := r.ParseForm(); err != nil {
		json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Unable to parse form data",
        })
		return
	}
	var post g.Post
	
	post.ID = uuid.New().String()
	post.Title = r.FormValue("title")
	post.Category = r.FormValue("myPostCategory")
	post.Content = r.FormValue("content")
	
	if err := insertPost(g.DB, post); err != nil {
		log.Println("Error inserting post:", err)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"error": "Failed to create post",
		})
		return
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "Post created successfully",
		"post":    post,
	})
}