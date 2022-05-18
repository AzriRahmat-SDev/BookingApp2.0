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

//Admin is a handler that parses a ResponseWriter and a Request from a POST.
//Within this handler it checks if an admin user has already logged in and also handles
//The removal of bookings and of staff members is done here
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

		//Concurrency is added here as to ensure that the booking are updated
		//if there are 2 users booking at the same time
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
	//Data to be passed to the template
	data["user"] = user
	data["doctorList"] = database.DoctorList

	if err := render.Template(w, r, "admin.page.html", &render.TemplateData{
		Data: data,
	}); err != nil {
		log.Println("Admin: Error parsing templates: ", err)
	}
}
