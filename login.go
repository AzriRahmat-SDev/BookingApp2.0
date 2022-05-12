package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func logIn(w http.ResponseWriter, r *http.Request) {
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
			Username: r.PostFormValue("username"),
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
			fmt.Println("Hello from line 38")
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
			fmt.Println("Hello from line 53")
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
			fmt.Println("Hello from line 66")
			log.Println("In user.isAdmin")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}
		//if the username and password is correct it will exe to the home page
		fmt.Println("Hello from line 72")
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
