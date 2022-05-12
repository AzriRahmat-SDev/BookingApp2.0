package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func admin(w http.ResponseWriter, r *http.Request) {

	data := make(map[string]interface{})

	user := getUser(r)

	if user == nil || !user.isAdmin {
		if err := Template(w, r, "restricted.page.html", &TemplateData{
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
			DeleteBookingFromBookingList(id)
			wg.Done()
		}(doctorId)

		go func(id int) {
			if err := recover(); err != nil {
				fmt.Println("Admin: ", err)
			}
			DeleteBookingFromBookingList(id)
			wg.Done()
		}(doctorId)

		if err := DeleteDoctor(doctorId); err != nil {
			log.Println("Admin: Error deleting venue ", err)
		}
		wg.Wait()
	}
	data["user"] = user
	data["doctorList"] = doctorList

	if err := Template(w, r, "admin.page.html", &TemplateData{
		Data: data,
	}); err != nil {
		log.Println("Admin: Error parsing templates: ", err)
	}
}
