package database

import (
	"fmt"
)

type Doctor struct {
	Id   int
	Name string
}

var DoctorList []*Doctor

func InitDoctors() {
	DoctorList = []*Doctor{
		{111, "Dr Idris"},
		{222, "Dr Vickram"},
		{333, "Dr Lim"},
	}
}
func incrementDoctor() int {
	max := 0
	//max := len(doctorList)
	for _, doctor := range DoctorList {
		if doctor.Id > max {
			max = doctor.Id
		}
	}
	return max + 1
}

func GetDoctorById(ID int) *Doctor {
	for i := 0; i < len(DoctorList); i++ {
		if DoctorList[i].Id == ID {
			return DoctorList[i]
		}
	}
	return nil
}

func AddDoctor(value *Doctor) {
	value.Id = incrementDoctor()
	DoctorList = append(DoctorList, value)
}

func DeleteDoctor(id int) error {
	for i, value := range DoctorList {
		if value.Id == id {
			DoctorList = append(DoctorList[:i], DoctorList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("index of error not found")
}
