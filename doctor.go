package main

import "fmt"

type Doctor struct {
	Id           int
	NameOfDoctor string
}

var doctorList []*Doctor

func initDoctors() {
	doctorList = []*Doctor{
		{1, "Dr Vickram"},
		{2, "Dr Idris"},
		{3, "Dr Lim"},
	}
}
func incrementDoctor() int {
	max := 0
	for _, doctor := range doctorList {
		if doctor.Id > max {
			max = doctor.Id
		}
	}
	return max + 1
}

func GetDoctorById(id int) *Doctor {
	for _, value := range doctorList {
		if value.Id == id {
			return value
		}
	}
	return nil
}

func addDoctor(value *Doctor) {
	value.Id = incrementDoctor()
	doctorList = append(doctorList, value)
}

func DeleteDoctor(id int) error {
	for i, value := range doctorList {
		if value.Id == id {
			doctorList = append(doctorList[:i], doctorList[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("index of error not found")
}
