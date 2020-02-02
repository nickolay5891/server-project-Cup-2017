package main

import "fmt"

//us.AllUsers[id] = User{"swe", "fggf", 6567, "32", 89, "hkj"}

//Функция для получения данных о пользователе
func (us *AllUsers) getUser(id uint32) (User, error) {
	var err error
	if _, ok := us.AllUsers[id]; ok {
		err = nil
	} else {
		err = fmt.Errorf("User is not found")
	}
	return us.AllUsers[id], err
}

//Функция для получения данных о достопримечательности
func (loc *AllLocations) getLocation(id uint32) (Location, error) {
	var err error
	if _, ok := loc.AllLocations[id]; ok {
		err = nil
	} else {
		err = fmt.Errorf("Location is not found")
	}
	return loc.AllLocations[id], err
}

//Функция для получения данных о посещении
func (loc *AllVisits) getVisit(id uint32) (Visit, error) {
	var err error
	if _, ok := loc.AllVisits[id]; ok {
		err = nil
	} else {
		err = fmt.Errorf("Visit is not found")
	}
	return loc.AllVisits[id], err
}
