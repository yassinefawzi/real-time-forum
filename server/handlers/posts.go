package handlers

import (
	"database/sql"
	"encoding/json"
	g "forum/server/global"
	"log"
	http "net/http"
	"github.com/google/uuid"
	session "forum/server/session"

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

func Getposts(w http.ResponseWriter, r *http.Request) {
    _, ok := session.GetSessionUsername(r)
    if !ok {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

    rows, err := g.DB.Query("SELECT title, content, category FROM posts ORDER BY rowid DESC")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var posts []g.Post
    for rows.Next() {
        var post g.Post
        if err := rows.Scan(&post.Title, &post.Content, &post.Category); err != nil {
            continue
        }
        posts = append(posts, post)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}