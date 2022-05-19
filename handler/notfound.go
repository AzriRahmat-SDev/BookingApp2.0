package handler

import (
	"GoInActionAssignment/database"
	"GoInActionAssignment/render"
	"log"
	"net/http"
)

//NotFound redirects the user to a 404 page
func NotFound(w http.ResponseWriter, r *http.Request) {
	var user database.User
	if o := getUser(r); o != nil {
		user = *o
	}
	data := make(map[string]interface{})
	data["user"] = user
	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		if err := render.Template(w, r, "notfound.page.html", &render.TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Notfound: Error parsing template: ", err)
		}
	}
}
