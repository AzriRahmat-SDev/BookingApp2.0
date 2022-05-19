package handler

import (
	"GoInActionAssignment/database"
	"GoInActionAssignment/form"
	"GoInActionAssignment/render"
	"log"
	"net/http"
)

//SignUp handles new user sign ups
//It creates a new user via variables such as firstname, lastname, username and password
//Input validation is used to ensure these field adhere to the conditions
//Current existing user are checked against the new user.
//Once passing the validation checks the handler will proceed to create a new user in the database
func SignUp(w http.ResponseWriter, r *http.Request) {

	if getUser(r) != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("Registering:", err)
		return
	}
	if r.Method == http.MethodPost {
		newUser := database.User{
			Firstname: r.FormValue("firstname"),
			Lastname:  r.FormValue("lastname"),
			Username:  r.FormValue("username"),
			Password:  []byte(r.FormValue("password")),
		}
		form := form.New(r.PostForm)
		form.Required("firstname", "lastname", "username", "password")

		if form.ExistingUser() {
			form.Errors.Add("username", "Username already in use")
		}

		if !form.Valid() {
			data := make(map[string]interface{})
			data["signup"] = newUser
			if err := render.Template(w, r, "signup.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Registration: ", err)
			}
			return
		}
		if err := database.CreateNewUser(&newUser); err != nil {
			log.Println("Registration: ", err)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := render.Template(w, r, "signup.page.html", &render.TemplateData{
		Data: make(map[string]interface{}),
		Form: form.New(nil)}); err != nil {
		log.Println("Registration: ", err)
		return
	}
}
