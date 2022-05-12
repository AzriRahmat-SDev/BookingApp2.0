package main

import (
	"log"
	"net/http"
	"strconv"
)

func adminBooking(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	user := getUser(r)

	if user == nil || !user.isAdmin {
		if err := Template(w, r, "restricted.page.html", &TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Admin: Error parsing template: ", err)
		}
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Register:", err)
		return
	}

	if r.Method == http.MethodPost {
		form := New(r.PostForm)
		newDoctor := Doctor{
			NameOfDoctor: r.FormValue("name"),
		}
		form.Required("name", "id")

		id, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			form.Errors.Add("id", "id only accepts intergers only")
		}

		if !form.Valid() {
			data := make(map[string]interface{})
			data["doctor"] = newDoctor
			data["user"] = user

			if err := Template(w, r, "adminbooking.page.html", &TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Admin: Venue:", err)
			}
			return
		}
		newDoctor.Id = id
		addDoctor(&newDoctor)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	if err := Template(w, r, "adminBooking.page.html",
		&TemplateData{
			Data: make(map[string]interface{}),
			Form: New(nil)}); err != nil {
		log.Println("Admin: Venue: ", err)
		return
	}
}
