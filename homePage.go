package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["doctors"] = doctorList

	user := getUser(r)
	if user != nil {
		data["user"] = user

		if r.Method == http.MethodPost {
			r.ParseForm()
			//checks Form Named <doctorBookingForm> as a valid form
			doctorId, err := strconv.Atoi(r.PostFormValue("doctorBookingForm"))
			if err != nil {
				log.Println("Home: ", err)
			}
			date := r.PostFormValue(fmt.Sprintf("date%d", doctorId))
			//checks validity of date selected with the current date
			if !bookingDateHandler(date) {
				form := New(r.Form)
				form.Errors.Add("date", fmt.Sprintf("Date selected has already passed! Please select another date"))
				if err := Template(w, r, "home.page.html", &TemplateData{Data: data, Form: form}); err != nil {
					log.Print("Home: ", err)
				}
				return
			}
			//checks user input date with respect of the doctors date of availability
			if !bookingIsAvail(doctorId, date) {
				form := New(r.Form)
				form.Errors.Add("date", fmt.Sprintf("Date selected for \"%s\" has already been booked! Please select another date", GetDoctorById(doctorId).NameOfDoctor))
				if err := Template(w, r, "home.page.html", &TemplateData{Data: data, Form: form}); err != nil {
					log.Print("Home: ", err)
				}
				return
			}

			newBookings := Booking{
				CustomerId: user.CustomerId,
				DoctorId:   doctorId,
				Date:       date,
			}
			bookingId := newBooking(newBookings)
			user.BookingId = append(user.BookingId, bookingId)
			form := New(r.Form)
			form.Errors.Add("success", fmt.Sprintf("Booking for \"%s\" on \"%s\" successful!", GetDoctorById(doctorId).NameOfDoctor, date))

			if err := Template(w, r, "home.page.html", &TemplateData{Data: data, Form: form}); err != nil {
				log.Println("Home: ", err)
			}
			return
		}
		if err := Template(w, r, "home.page.html", &TemplateData{Data: data, Form: New(nil)}); err != nil {
			log.Println("Home: ", err)
		}
	} else {
		if err := Template(w, r, "home.page.html",
			&TemplateData{
				Data: data,
				Form: New(nil)}); err != nil {
			log.Println("Home: ", err)

		}
	}
}

func getUser(r *http.Request) (user *User) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}

	if username, ok := Sessions[cookie.Value]; ok {
		user = Users[username]
	}
	return
}
