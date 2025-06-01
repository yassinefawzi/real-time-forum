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
	"strings"
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
    // Regular form parsing works now
    if err := r.ParseForm(); err != nil {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Unable to parse form data",
        })
        return
    }
    
    var user g.User
    
    user.ID = uuid.New().String()
    user.Username = r.FormValue("Nickname")
    user.Email = r.FormValue("E-mail")
    
    // Debug: Print received form values
    log.Printf("Received form values:")
    log.Printf("Nickname: '%s'", r.FormValue("Nickname"))
    log.Printf("E-mail: '%s'", r.FormValue("E-mail"))
    log.Printf("Age: '%s'", r.FormValue("Age"))
    
    // Validate required fields
    if user.Username == "" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Nickname is required",
        })
        return
    }
    
    if user.Email == "" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Email is required",
        })
        return
    }
    
    age, err := strconv.Atoi(r.FormValue("Age"))
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Invalid age value",
        })
        return
    }
    user.Age = age
    
    user.Gender = r.FormValue("gender")
    user.FirstName = r.FormValue("First Name")
    user.LastName = r.FormValue("Last Name")
    
    password := r.FormValue("Password")
    if password == "" {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Password is required",
        })
        return
    }
    
    user.Password, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": "Error processing password",
        })
        return
    }
    
    if err := insertUser(g.DB, user); err != nil {
        log.Println("Insert error:", err)
        
        w.Header().Set("Content-Type", "application/json")
        
        errorMsg := "Registration failed. Please try again."
        errorStr := err.Error()
        
        if strings.Contains(errorStr, "UNIQUE constraint failed: users.email") {
            errorMsg = "Email already exists. Please use a different email."
        } else if strings.Contains(errorStr, "UNIQUE constraint failed: users.username") {
            errorMsg = "Username already exists. Please choose a different username."
        }
        
        json.NewEncoder(w).Encode(map[string]interface{}{
            "success": false,
            "error": errorMsg,
        })
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "success": true,
        "message": "Registration successful!",
    })
}
