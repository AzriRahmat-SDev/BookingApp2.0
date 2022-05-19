package handler

import (
	"GoInActionAssignment/database"
	"GoInActionAssignment/form"
	"GoInActionAssignment/render"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//function logIn is a type of Handler(due to how parameters are written) that can be used in ListenAndServe as it accepts a handler type
//http. denotes the call from the http package "net/http"
//Therefore w of type http.ResponseWriter and r of type http.Request
func LogIn(w http.ResponseWriter, r *http.Request) {

	//The convention of declaring != nil means that when the getUser() func is executed
	//there is a respose hence redirecting back to root
	if getUser(r) != nil {
		log.Println("Resuming previous session")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	//err is declared as a parseForm method, it is executed in the back and returns a result
	//of which if error is not empty, it will log the err
	if err := r.ParseForm(); err != nil {
		log.Println("Login:", err)
		return
	}

	//Check if the Method is a Post then runs the code
	if r.Method == http.MethodPost {
		user := database.User{
			Username: r.PostFormValue("username"),
			Password: []byte(r.PostFormValue("password")),
		}

		form := form.New(r.PostForm)
		form.Required("username", "password")

		//If user Exist(return value is false), it will show an error message
		if !form.ExistingUser() {
			form.Errors.Add("username", "Username and/or password do not match")
		} else {
			//Using bcrypt to compare the user input with the hash password in the database
			if err := bcrypt.CompareHashAndPassword(database.Users[user.Username].Password, user.Password); err != nil {
				form.Errors.Add("username", "Username and/or password do not match")
			}
		}
		//Goes into the input validation fields and check for any Errors
		//Input validation is preferably used before the data is used for computing
		if !form.Valid() {
			log.Println("Form is not valid")
			data := make(map[string]interface{})
			data["login"] = user
			if err := render.Template(w, r, "login.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Login: ", err)
			}
			return
		}

		//Generating random UUID for Session management and cookies
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
		database.Sessions[cookie.Value] = user.Username
		u := database.Users[user.Username]

		//logging if admin or user has been sign in
		if u.IsAdmin {
			log.Println("Successful Admin login")
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		} else {
			log.Println("Successful User login")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	if err := render.Template(w, r, "login.page.html",
		&render.TemplateData{
			Data: make(map[string]interface{}),
			Form: form.New(nil)}); err != nil {
		log.Println("Login: ", err)
		return
	}
}
