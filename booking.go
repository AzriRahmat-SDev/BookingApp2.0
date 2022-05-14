package main

import "time"

type Booking struct {
	DoctorId   int
	CustomerId int
	BookingId  int
	Date       string
}

var BookingList map[int]Booking
var bookingId int

func init() {
	BookingList = make(map[int]Booking)
	list := []Booking{
		{1, 1, 1, "2022-06-06"},
		{2, 1, 1, "2022-11-07"},
		{3, 1, 1, "2022-08-31"},
	}

	for _, value := range list {
		newBooking(value)
	}

}

func bookingIsAvail(doctorId int, date string) bool {
	for _, value := range BookingList {
		if value.DoctorId == doctorId && value.Date == date {
			return false
		}
	}
	return true
}

func newBooking(value Booking) int {
	bookingId++
	value.BookingId = bookingId
	BookingList[bookingId] = value
	return bookingId
}

func (b *Booking) GetDoctorName() string {
	return GetDoctorById(b.BookingId).NameOfDoctor
}

func DeleteBookingFromBookingList(id int) error {
	for result, value := range BookingList {
		if value.DoctorId == id {
			delete(BookingList, result)
		}
	}
	return nil
}

func bookingDateHandler(date string) bool {
	currentDate := time.Now()
	currentDateString := currentDate.String()
	if currentDateString > date {
		return false
	}
	return true
}
