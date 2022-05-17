package handler

import (
	"GoInActionAssignment/internal/database"
	"fmt"
	"net/http"
	"time"
)

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
