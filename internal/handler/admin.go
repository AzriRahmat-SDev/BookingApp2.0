package handler

import (
	"GoInActionAssignment/internal/database"
	"GoInActionAssignment/internal/render"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func Admin(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	user := getUser(r)

	if user == nil || !user.IsAdmin {
		if err := render.Template(w, r, "restricted.page.html", &render.TemplateData{
			Data: data,
		}); err != nil {
			log.Println("Admin: Erroring parsing template: ", err)
		}
		return
	}
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("Admin: ", err)
		}
		doctorId, err := strconv.Atoi(r.PostFormValue("doctorId"))
		if err != nil {
			log.Println("Admin: parsing string to int: ", err)
		}

		var wg sync.WaitGroup
		wg.Add(2)

		go func(id int) {
			if err := recover(); err != nil {
				fmt.Println("Admin: ", err)
			}
			database.DeleteBookingFromBookingList(id)
			wg.Done()
		}(doctorId)

		go func(id int) {
			if err := recover(); err != nil {
				fmt.Println("Admin: ", err)
			}
			database.DeleteBookingFromBookingList(id)
			wg.Done()
		}(doctorId)

		if err := database.DeleteDoctor(doctorId); err != nil {
			log.Println("Admin: Error deleting venue ", err)
		}
		wg.Wait()
	}
	data["user"] = user
	data["doctorListAdmin"] = database.DoctorList

	if err := render.Template(w, r, "admin.page.html", &render.TemplateData{
		Data: data,
	}); err != nil {
		log.Println("Admin: Error parsing templates: ", err)
	}
}
