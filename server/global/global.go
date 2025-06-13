package glo
import (
	"database/sql"
)

var DB *sql.DB

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
	Title       string
	Content     string
	Category    string
}