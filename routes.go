package main

import (
	"net/http"
)

func routes() http.Handler {

	mux := http.NewServeMux()

	//Handlefunc(pattern, handler)
	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/login", logIn)
	mux.HandleFunc("/signup", signUp)
	mux.HandleFunc("/logout", logOutUser)
	mux.HandleFunc("/bookings", currentBookings)
	mux.HandleFunc("/admin", admin)
	mux.HandleFunc("/adminBooking", adminBooking)
	mux.HandleFunc("/", notFound)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	return mux
}
