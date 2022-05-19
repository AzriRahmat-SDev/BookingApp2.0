package main

import (
	"GoInActionAssignment/handler"
	"net/http"
)

func routes() http.Handler {

	mux := http.NewServeMux()

	//Handlefunc(pattern, handler)
	mux.HandleFunc("/home", handler.HomePage)
	mux.HandleFunc("/login", handler.LogIn)
	mux.HandleFunc("/signup", handler.SignUp)
	mux.HandleFunc("/logout", handler.LogOutUser)
	mux.HandleFunc("/bookings", handler.CurrentBookings)
	mux.HandleFunc("/admin", handler.Admin)
	mux.HandleFunc("/addDoctor", handler.AddDoctor)
	mux.HandleFunc("/", handler.NotFound)

	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	return mux
}
