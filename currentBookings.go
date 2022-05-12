package main

import (
	"log"
	"net/http"
	"strconv"
)

func currentBookings(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	user := getUser(r)

	if user == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	data["user"] = user

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			log.Println("Bookings: ", err)
		}
		bookingID, err := strconv.Atoi(r.PostFormValue("bookingID"))
		if err != nil {
			log.Println("Booking: Parsing String to int: ", err)
		}
		if err := user.cancelBookings(bookingID); err != nil {
			log.Fatalln(err)
		}

		delete(BookingList, bookingID)
	}
	usersBookings := []Booking{}

	for _, v := range user.BookingId {
		usersBookings = append(usersBookings, BookingList[v])
	}
	data["bookinglist"] = usersBookings

	if err := Template(w, r, "bookings.page.html", &TemplateData{Data: data, Form: New(nil)}); err != nil {
		log.Println("Bookings: ", err)
	}
}
