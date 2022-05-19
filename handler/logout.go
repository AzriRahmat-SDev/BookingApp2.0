package handler

import (
	"GoInActionAssignment/database"
	"fmt"
	"net/http"
	"time"
)

//LogOutUser parse a writer and a response a POST request
//Within this handler i check if a user is logged
//Deletes the current seesion and cookies
//Logs the time since the session was and reroutes the user to the home page
func LogOutUser(w http.ResponseWriter, r *http.Request) {
	if getUser(r) == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	c, _ := r.Cookie("session")
	start := time.Now()

	delete(database.Sessions, c.Value)

	c = &http.Cookie{
		Name:   "session",
		MaxAge: -1,
		Value:  "",
	}
	http.SetCookie(w, c)

	fmt.Println(time.Since(start))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
