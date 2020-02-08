package main

import (
	"fmt"
	"sort"
)

//us.AllUsers[id] = User{"swe", "fggf", 6567, "32", 89, "hkj"}

//Функция для получения данных о пользователе
func (us *AllUsers) getUser(id uint32) (User, error) {
	var err error
	if _, ok := us.AllUsers[id]; ok {
		err = nil
	} else {
		err = fmt.Errorf("User is not found")
	}
	//return nil, err
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
func (vis *AllVisits) getVisit(id uint32) (Visit, error) {
	var err error
	if _, ok := vis.AllVisits[id]; ok {
		err = nil
	} else {
		err = fmt.Errorf("Visit is not found")
	}
	return vis.AllVisits[id], err
}

//UserVisit Структура посещения пользователя
type UserVisit struct {
	Mark      uint32
	VisitedAt int
	Place     string
}

//UserVisits Все посещения пользователя
type UserVisits struct {
	Visits []UserVisit
}

//Функция для получения списка мест, которые посетил пользователь
func getUserVisits(userID uint32, fromDate *int, toDate *int, country *string, toDistance *uint32) (*UserVisits, error) {
	if _, err := (&allUsers).getUser(userID); err != nil {
		return nil, err
	}
	var userVis []UserVisit
	for _, visit := range allVisits.AllVisits {
		if visit.User != userID {
			continue
		}
		if fromDate != nil && visit.VisitedAt < *fromDate {
			continue
		}
		if toDate != nil && visit.VisitedAt > *toDate {
			continue
		}
		if country != nil && allLocations.AllLocations[visit.Location].Country != *country {
			continue
		}
		if toDistance != nil && allLocations.AllLocations[visit.Location].Distance >= *toDistance {
			continue
		}
		userVisit := UserVisit{visit.Mark, visit.VisitedAt, allLocations.AllLocations[visit.Location].Place}
		userVis = append(userVis, userVisit)
	}
	sort.SliceStable(userVis, func(i, j int) bool { return userVis[i].VisitedAt < userVis[j].VisitedAt })
	a := UserVisits{userVis}
	return &a, nil
}

//Функция для получения средней оценки достопримечательности
func getLocationsAvg(locID uint32, fromDate *int, toDate *int, fromAge *uint32, toAge *uint32, gender *string) (avg float64, err error) {
	if _, err := (&allLocations).getLocation(locID); err != nil {
		return avg, err
	}
	var sumMarks uint32
	var marks []uint32
	for _, visit := range allVisits.AllVisits {
		if visit.Location != locID {
			continue
		}
		if fromDate != nil && visit.VisitedAt < *fromDate {
			continue
		}
		if toDate != nil && visit.VisitedAt > *toDate {
			continue
		}
		if fromAge != nil && getAgeUser(allUsers.AllUsers[visit.User].BirthDate) < *fromAge {
			continue
		}
		if toAge != nil && getAgeUser(allUsers.AllUsers[visit.User].BirthDate) > *toAge {
			continue
		}
		if gender != nil && allUsers.AllUsers[visit.User].Gender != *gender {
			continue
		}
		sumMarks += visit.Mark
		marks = append(marks, visit.Mark)
	}
	if len(marks) > 0 {
		avg = round(float64(sumMarks) / float64(len(marks)))
	}
	return avg, nil
}
