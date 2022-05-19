package handler

import (
	"GoInActionAssignment/database"
	"GoInActionAssignment/form"
	"GoInActionAssignment/render"
	"log"
	"net/http"
	"strconv"
)

func CurrentBookings(w http.ResponseWriter, r *http.Request) {
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
		bookingId, err := strconv.Atoi(r.PostFormValue("bookingId"))
		if err != nil {
			log.Println("Booking: Parsing String to int: ", err)
		}
		if err := user.CancelBookings(bookingId); err != nil {
			log.Fatalln(err)
		}

		delete(database.BookingList, bookingId)
	}
	usersBookings := []database.Booking{}

	for _, v := range user.BookingId {
		usersBookings = append(usersBookings, database.BookingList[v])
	}
	data["bookingList"] = usersBookings
	//log.Println(data)
	if err := render.Template(w, r, "bookings.page.html", &render.TemplateData{Data: data, Form: form.New(nil)}); err != nil {
		log.Println("Bookings: ", err)
	}
}
