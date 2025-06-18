package session

import (
	"crypto/rand"
	"encoding/hex"
    "log"
	"net/http"
    "time"
	g "forum/server/global"
)

func CreateSession(userID string) string {
	id := make([]byte, 16)
	if _, err := rand.Read(id); err != nil {
        log.Fatal("Failed to generate session ID:", err)
	}
	return hex.EncodeToString(id)
}

func SetSession(w http.ResponseWriter, username string) {
	sessionId := CreateSession(username)
	g.SessionsMu.Lock()
	g.Sessions[sessionId] = username
	g.SessionsMu.Unlock()

	cookie := &http.Cookie{
		Name:    "session_id",
		Value:  sessionId,
		HttpOnly: true,
		Path: "/",
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
}

func GetSessionUsername(r *http.Request) (string, bool) {
	cookie, err := r.Cookie("session_id")
	if err != nil {
		return "", false
	}
	g.SessionsMu.Lock()
	username, exists := g.Sessions[cookie.Value], true
	g.SessionsMu.Unlock()

	return username, exists
}

func DeleteSession(w http.ResponseWriter, r *http.Request) {
	log.Println("Deleting session")

	cookie, err := r.Cookie("session_id")
	if err != nil {
		return
	}

	g.SessionsMu.Lock()
	delete(g.Sessions, cookie.Value)
	g.SessionsMu.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:     "session_id",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	})
}
