package main

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/microcosm-cc/bluemonday"
	"golang.org/x/crypto/bcrypt"
)

//function logIn is a type of Handler(due to how parameters are written) that can be used in ListenAndServe as it accepts a handler type
//http. denotes the call from the http package "net/http"
//Therefore w of type http.ResponseWriter and r of type http.Request
func logIn(w http.ResponseWriter, r *http.Request) {

	p := bluemonday.UGCPolicy()

	if getUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Login:", err)
		return
	}

	if r.Method == http.MethodPost {
		user := User{
			Username: p.Sanitize(r.PostFormValue("username")),
			Password: []byte(r.PostFormValue("password")),
		}

		form := New(r.PostForm)
		form.Required("username", "password")

		if !form.ExistingUser() {
			form.Errors.Add("username", "Username and/or password do not match")
		} else {
			if err := bcrypt.CompareHashAndPassword(Users[user.Username].Password, user.Password); err != nil {
				form.Errors.Add("username", "Username and/or password do not match")
			}
		}
		if !form.Valid() {
			data := make(map[string]interface{})
			data["login"] = user
			if err := Template(w, r, "login.page.html", &TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Login: ", err)
			}
			return
		}

		id, err := uuid.NewRandom()
		if err != nil {
			log.Println("Login:", err)
		}
		cookie := &http.Cookie{
			Name:     "session",
			Value:    id.String(),
			HttpOnly: true,
		}
		http.SetCookie(w, cookie)
		Sessions[cookie.Value] = user.Username
		u := Users[user.Username]

		if u.isAdmin {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if err := Template(w, r, "login.page.html",
		&TemplateData{
			Data: make(map[string]interface{}),
			Form: New(nil)}); err != nil {
		log.Println("Login: ", err)
		return
	}
}
