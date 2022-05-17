package handler

import (
	"GoInActionAssignment/internal/database"
	"GoInActionAssignment/internal/form"
	"GoInActionAssignment/internal/render"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["doctors"] = database.DoctorList

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
			if !database.BookingDateHandler(date) {
				form := form.New(r.Form)
				form.Errors.Add("date", fmt.Sprintf("Date selected has already passed! Please select another date"))
				if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form}); err != nil {
					log.Print("Home: ", err)
				}
				return
			}
			//checks user input date with respect of the doctors date of availability
			if !database.BookingIsAvail(doctorId, date) {
				form := form.New(r.Form)
				form.Errors.Add("date", fmt.Sprintf("Date selected for \"%s\" has already been booked! Please select another date", database.GetDoctorById(doctorId).Name))
				if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form}); err != nil {
					log.Print("Home: ", err)
				}
				return
			}

			newBookings := database.Booking{
				CustomerId: user.CustomerId,
				DoctorId:   doctorId,
				Date:       date,
			}
			bookingId := database.NewBooking(newBookings)
			user.BookingId = append(user.BookingId, bookingId)
			form := form.New(r.Form)
			form.Errors.Add("success", fmt.Sprintf("Booking for \"%s\" on \"%s\" successful!", database.GetDoctorById(doctorId).Name, date))

			if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form}); err != nil {
				log.Println("Home: ", err)
			}
			return
		}
		if err := render.Template(w, r, "home.page.html", &render.TemplateData{Data: data, Form: form.New(nil)}); err != nil {
			log.Println("Home: ", err)
		}
	} else {
		if err := render.Template(w, r, "home.page.html",
			&render.TemplateData{
				Data: data,
				Form: form.New(nil)}); err != nil {
			log.Println("Home: ", err)

		}
	}
}

func getUser(r *http.Request) (user *database.User) {
	cookie, err := r.Cookie("session")
	if err != nil {
		return nil
	}

	if username, ok := database.Sessions[cookie.Value]; ok {
		user = database.Users[username]
	}
	return
}
