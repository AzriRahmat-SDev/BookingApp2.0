//Package database provide the necessary initialization of data,
//adding, deleting and input validation checks needed by the handler functions
//in the handler package.
package database

import "time"

//This File Provide the necessary functions to carry out
//Initialization,adding, deleting of available bookings needed from the handler functions
type Booking struct {
	BookingId  int
	CustomerId int
	DoctorId   int
	Date       string
}

var BookingList map[int]Booking
var BookingId int

//Initialize current existing bookings
func init() {
	BookingList = make(map[int]Booking)
	list := []Booking{
		{1, 1, 111, "2021-03-31"},
		{2, 1, 222, "2021-04-03"},
		{3, 1, 333, "2021-04-02"},
	}

	for _, value := range list {
		NewBooking(value)
	}

}

//Checks if the booking is available on a selected day
func BookingIsAvail(doctorId int, date string) bool {
	for _, value := range BookingList {
		if value.DoctorId == doctorId && value.Date == date {
			return false
		}
	}
	return true
}

//Add new bookings
func NewBooking(value Booking) int {
	BookingId++
	value.BookingId = BookingId
	BookingList[BookingId] = value
	return BookingId
}

//Get Doctor by name
func (b *Booking) GetDoctorName() string {
	return GetDoctorById(b.DoctorId).Name
}

func DeleteBookingFromBookingList(id int) error {
	for result, value := range BookingList {
		if value.DoctorId == id {
			delete(BookingList, result)
		}
	}
	return nil
}

func BookingDateHandler(date string) bool {
	currentDate := time.Now()
	currentDateString := currentDate.String()
	if currentDateString > date {
		return false
	}
	return true
}
