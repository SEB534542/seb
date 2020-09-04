package seb

import (
	"net/http"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var DbSessions = map[string]string{} // session ID, user ID
var DbUsers = map[string]user{}      // user ID, user

// AlreadyLoggedIn checks if the UUID in the "session" cookie
// exists in the variable DbSessions and if that username still exist in our DbUsers
func AlreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := DbSessions[c.Value]
	_, ok := DbUsers[un]
	return ok
}

// GetIP gets a requests IP address by reading off the forwarded-for
// header (for proxies) and falls back to use the remote address.
func GetIP(req *http.Request) string {
	forwarded := req.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return req.RemoteAddr
}
