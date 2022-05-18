package handler

import (
	"GoInActionAssignment/internal/database"
	"GoInActionAssignment/internal/form"
	"GoInActionAssignment/internal/render"
	"log"
	"net/http"
	"strconv"
)

//AddDoctor is a handler that parses a ResponseWriter and a Request from a POST.
//Checks if admin user has already logged in
//Addition of staff members are done here
func AddDoctor(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})

	user := getUser(r)

	if user == nil || !user.IsAdmin {
		if err := render.Template(w, r, "restricted.page.html", &render.TemplateData{
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
		form := form.New(r.PostForm)
		newDoctor := database.Doctor{
			Name: r.FormValue("name"),
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

			if err := render.Template(w, r, "addDoctor.page.html", &render.TemplateData{
				Data: data,
				Form: form,
			}); err != nil {
				log.Println("Admin: Venue:", err)
			}
			return
		}
		newDoctor.Id = id
		database.AddDoctor(&newDoctor)
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	if err := render.Template(w, r, "addDoctor.page.html",
		&render.TemplateData{
			Data: make(map[string]interface{}),
			Form: form.New(nil)}); err != nil {
		log.Println("Admin: Venue: ", err)
		return
	}
}
