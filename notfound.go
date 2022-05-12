package main

import (
	"log"
	"net/http"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	var user User
	if o := getUser(r); o != nil {
		user = *o
	}
	data := make(map[string]interface{})
	data["user"] = user
	switch r.URL.Path {
	case "/":
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	default:
		if err := Template(w, r, "notfound.page.html", &TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Notfound: Error parsing template: ", err)
		}
	}
}
