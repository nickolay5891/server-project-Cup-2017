package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

//Функция для получения списка JSON файлов
func getListOfFiles() (listOfFiles []string) {
	files, err := ioutil.ReadDir("data")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			listOfFiles = append(listOfFiles, file.Name())
		}
	}
	return
}

//Создание структур
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int    `json:"birth_date"`
	Gender    string `json:"gender"`
	ID        uint32 `json:"id"`
	Email     string `json:"email"`
}

type Users struct {
	Users []User
}

type AllUsers struct {
	AllUsers map[uint32]User
}

type Location struct {
	Distance uint32 `json:"distance"`
	City     string `json:"city"`
	Place    string `json:"place"`
	ID       uint32 `json:"id"`
	Country  string `json:"country"`
}

type Locations struct {
	Locations []Location
}

type AllLocations struct {
	AllLocations map[uint32]Location
}

type Visit struct {
	User      uint32 `json:"user"`
	Location  uint32 `json:"location"`
	VisitedAt int    `json:"visited_at"`
	ID        uint32 `json:"id"`
	Mark      uint32 `json:"mark"`
}

type Visits struct {
	Visits []Visit
}

type AllVisits struct {
	AllVisits map[uint32]Visit
}

//Функции для извлечения данных из JSON-файлов и их преобразования
func getDataUsers(fileName string) Users {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataUsers Users
	json.Unmarshal(file, &dataUsers)
	return dataUsers
}

func getDataLocations(fileName string) Locations {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataLocations Locations
	json.Unmarshal(file, &dataLocations)
	return dataLocations
}

func getDataVisits(fileName string) Visits {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataVisits Visits
	json.Unmarshal(file, &dataVisits)
	return dataVisits
}

//Функция для создания БД allUsers, allLocations, allVisits
func getBasesData(listOfFiles []string) (AllUsers, AllLocations, AllVisits) {
	allUsers := map[uint32]User{}
	allLocations := map[uint32]Location{}
	allVisits := map[uint32]Visit{}
	for _, fileName := range listOfFiles {
		if strings.Contains(fileName, "users") {
			dataUsers := getDataUsers("data/" + fileName)
			for _, dataUser := range dataUsers.Users {
				allUsers[dataUser.ID] = dataUser
			}
		} else if strings.Contains(fileName, "locations") {
			dataLocations := getDataLocations("data/" + fileName)
			for _, dataLocation := range dataLocations.Locations {
				allLocations[dataLocation.ID] = dataLocation
			}
		} else if strings.Contains(fileName, "visits") {
			dataVisits := getDataVisits("data/" + fileName)
			for _, dataVisit := range dataVisits.Visits {
				allVisits[dataVisit.ID] = dataVisit
			}
		}
	}
	allUs := AllUsers{allUsers}
	allLoc := AllLocations{allLocations}
	allVis := AllVisits{allVisits}
	return allUs, allLoc, allVis
}
