package handlers

import (
	http "net/http"
	g "forum/server/global"
	"log"
	_ "github.com/mattn/go-sqlite3"
	"strconv"
	"database/sql"
	"golang.org/x/crypto/bcrypt"
	"encoding/json"
	"github.com/google/uuid"
)

func insertUser(db *sql.DB, user g.User) error {
    query := `
        INSERT INTO users (
            id, username, email, age, gender, firstName, lastName, password_hash
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?);
    `

    _, err := db.Exec(query, user.ID, user.Username, user.Email, user.Age, user.Gender, user.FirstName, user.LastName, user.Password,
    )

    return err
}


func Getregister(w http.ResponseWriter, r *http.Request) {
	var user g.User
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}
	user.ID = uuid.New().String()
	user.Username = r.FormValue("Nickname")
	user.Email = r.FormValue("E-mail")
	user.Age, _ = strconv.Atoi(r.FormValue("Age"))
	user.Gender = r.FormValue("gender")
	user.FirstName = r.FormValue("First Name")
	user.LastName = r.FormValue("Last Name")
	user.Password, _ = bcrypt.GenerateFromPassword([]byte(r.FormValue("Password")), bcrypt.DefaultCost)
	if err := insertUser(g.DB, user); err != nil {
        log.Println("Insert error:", err)
		http.Redirect(w, r, "/api/signup", http.StatusSeeOther)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func CheckUserExists(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	username := r.URL.Query().Get("username")
	var emailExists, usernameExists bool

	err := g.DB.QueryRow(`
		SELECT 
			EXISTS(SELECT 1 FROM users WHERE email = ?),
			EXISTS(SELECT 1 FROM users WHERE username = ?)
	`, email, username).Scan(&emailExists, &usernameExists)
	if err != nil {
		log.Println("Error checking user existence:", err)
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
		json.NewEncoder(w).Encode(map[string]bool{
		"emailExists":    emailExists,
		"usernameExists": usernameExists,
	})
}