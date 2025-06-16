package handlers

import (
	"database/sql"
	"encoding/json"
	"strings"
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

    rows, err := g.DB.Query("SELECT id, title, content, category FROM posts ORDER BY rowid DESC")
    if err != nil {
        http.Error(w, "Database error", http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var posts []g.Post
    for rows.Next() {
        var post g.Post
        if err := rows.Scan(&post.ID, &post.Title, &post.Content, &post.Category); err != nil {
            continue
        }
        posts = append(posts, post)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(posts)
}

func GetSinglePost(w http.ResponseWriter, r *http.Request) {
    // Extract the post ID from the URL manually
    pathParts := strings.Split(r.URL.Path, "/")
    if len(pathParts) < 4 || pathParts[3] == "" {
        http.Error(w, "Post ID not specified", http.StatusBadRequest)
        return
    }
    id := pathParts[3] // e.g. /api/singlepost/{id}

    // Prepare query
    var post g.Post
    err := g.DB.QueryRow("SELECT id, title, content, category FROM posts WHERE id = ?", id).
        Scan(&post.ID, &post.Title, &post.Content, &post.Category)

    if err != nil {
        if err == sql.ErrNoRows {
            http.Error(w, "Post not found", http.StatusNotFound)
        } else {
            http.Error(w, "Database error", http.StatusInternalServerError)
        }
        return
    }

    // Return JSON response
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(post)
}