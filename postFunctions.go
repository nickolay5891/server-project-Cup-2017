package main

import (
	"fmt"
	"strconv"
)

//Создание структур "тело запросов"
type UserPost struct {
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	BirthDate *int    `json:"birth_date"`
	Gender    *string `json:"gender"`
	Email     *string `json:"email"`
}

type LocationPost struct {
	Distance *uint32 `json:"distance"`
	City     *string `json:"city"`
	Place    *string `json:"place"`
	Country  *string `json:"country"`
}

type VisitPost struct {
	User      *uint32 `json:"user"`
	Location  *uint32 `json:"location"`
	VisitedAt *int    `json:"visited_at"`
	Mark      *uint32 `json:"mark"`
}

//Функция для обновления данных о пользователе
func (au *AllUsers) updateUser(id uint32, up UserPost) (err error) {
	if _, ok := au.AllUsers[id]; !ok {
		err = fmt.Errorf("User is not found")
		return
	}
	upUs := au.AllUsers[id]
	if up.FirstName != nil {
		upUs.FirstName = *up.FirstName
	}
	if up.LastName != nil {
		upUs.LastName = *up.LastName
	}
	if up.BirthDate != nil {
		upUs.BirthDate = *up.BirthDate
	}
	if up.Gender != nil {
		upUs.Gender = *up.Gender
	}
	if up.Email != nil {
		upUs.Email = *up.Email
	}
	au.AllUsers[id] = upUs
	return nil
}

//Функция для добавления нового пользователя пользователе
func (au *AllUsers) createUser(u User) (err error) {
	if _, ok := au.AllUsers[u.ID]; ok {
		err = fmt.Errorf("User with id = " + strconv.Itoa(int(u.ID)) + " exists")
		return
	}
	au.AllUsers[u.ID] = u
	return nil
}

//Функция для обновления данных о достопримечательности
func (al *AllLocations) updateLocation(id uint32, lp LocationPost) (err error) {
	if _, ok := al.AllLocations[id]; !ok {
		err = fmt.Errorf("Location is not found")
		return
	}
	upLoc := al.AllLocations[id]
	if lp.Distance != nil {
		upLoc.Distance = *lp.Distance
	}
	if lp.City != nil {
		upLoc.City = *lp.City
	}
	if lp.Place != nil {
		upLoc.Place = *lp.Place
	}
	if lp.Country != nil {
		upLoc.Country = *lp.Country
	}
	al.AllLocations[id] = upLoc
	return nil
}

//Функция для добавления новой достопримечательности
func (al *AllLocations) createLocation(l Location) (err error) {
	if _, ok := al.AllLocations[l.ID]; ok {
		err = fmt.Errorf("Location with id = " + strconv.Itoa(int(l.ID)) + " exists")
		return
	}
	al.AllLocations[l.ID] = l
	return nil
}

//Функция для обновления данных о посещении
func (av *AllVisits) updateVisit(id uint32, vp VisitPost) (err error) {
	if _, ok := av.AllVisits[id]; !ok {
		err = fmt.Errorf("Visit is not found")
		return
	}
	upVis := av.AllVisits[id]
	if vp.User != nil {
		upVis.User = *vp.User
	}
	if vp.Location != nil {
		upVis.Location = *vp.Location
	}
	if vp.VisitedAt != nil {
		upVis.VisitedAt = *vp.VisitedAt
	}
	if vp.Mark != nil {
		upVis.Mark = *vp.Mark
	}
	av.AllVisits[id] = upVis
	return nil
}

//Функция для добавления нового посещения
func (av *AllVisits) createVisit(v Visit) (err error) {
	if _, ok := av.AllVisits[v.ID]; ok {
		err = fmt.Errorf("Visit with id = " + strconv.Itoa(int(v.ID)) + " exists")
		return
	}
	av.AllVisits[v.ID] = v
	return nil
}
