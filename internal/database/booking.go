package database

import "time"

type Booking struct {
	BookingId  int
	CustomerId int
	DoctorId   int
	Date       string
}

var BookingList map[int]Booking
var BookingId int

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

func BookingIsAvail(doctorId int, date string) bool {
	for _, value := range BookingList {
		if value.DoctorId == doctorId && value.Date == date {
			return false
		}
	}
	return true
}

func NewBooking(value Booking) int {
	BookingId++
	value.BookingId = BookingId
	BookingList[BookingId] = value
	return BookingId
}

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
