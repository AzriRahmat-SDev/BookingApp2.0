package main

import (
	"log"
	"net/http"
)

func signUp(w http.ResponseWriter, r *http.Request) {

	if getUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Registering:", err)
		return
	}
	if r.Method == http.MethodPost {
		newUser := User{
			Firstname: r.FormValue("firstname"),
			Lastname:  r.FormValue("lastname"),
			Username:  r.FormValue("username"),
			Password:  []byte(r.FormValue("password")),
		}
		form := New(r.PostForm)

		form.Required("firstname", "lastname", "username", "password")
		if !form.Valid() {
			data := make(map[string]interface{})
			data["register"] = newUser

			if err := Template(w, r, "signup.page.html", &TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Registration: ", err)
			}
			return
		}
		if err := CreateNewUser(&newUser); err != nil {
			log.Println("Registration: ", err)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := Template(w, r, "signup.page.html", &TemplateData{
		Data: make(map[string]interface{}),
		Form: New(nil)}); err != nil {
		log.Println("Registration: ", err)
		return
	}
}
