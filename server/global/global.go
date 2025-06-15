package glo
import (
	"sync"
	"database/sql"
)

var DB *sql.DB

var (
    Sessions   = map[string]string{}
    SessionsMu sync.Mutex
)

type User struct {
	ID           string    
	Username     string    
	Email        string    
	Age          int       
	Gender       string    
	FirstName    string    
	LastName     string    
	Password	 []byte    
}

type Post struct {
	ID		  	string
	Title       string `json:"title"`
	Content     string `json:"content"`
	Category    string `json:"category"`
}