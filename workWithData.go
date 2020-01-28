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

//Создание структур users, locations, visits
type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int    `json:"birth_date"`
	Gender    string
	ID        uint32
	Email     string
}

type users struct {
	Users []user
}

type location struct {
	Distance uint32
	City     string
	Place    string
	ID       uint32
	country  string
}

type locations struct {
	Locations []location
}

type visit struct {
	User      uint32
	Location  uint32
	VisitedAt int `json:"visited_at"`
	ID        uint32
	Mark      uint32
}

type visits struct {
	Visits []visit
}

//Функции для извлечения данных из JSON-файлов и их преобразования
func getDataUsers(fileName string) users {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataUsers users
	json.Unmarshal(file, &dataUsers)
	return dataUsers
}

func getDataLocations(fileName string) locations {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataLocations locations
	json.Unmarshal(file, &dataLocations)
	return dataLocations
}

func getDataVisits(fileName string) visits {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	var dataVisits visits
	json.Unmarshal(file, &dataVisits)
	return dataVisits
}

//Функция для создания БД allUsers, allLocations, allVisits
func getBasesData(listOfFiles []string) (map[uint32]user, map[uint32]location, map[uint32]visit) {
	allUsers := map[uint32]user{}
	allLocations := map[uint32]location{}
	allVisits := map[uint32]visit{}
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
	return allUsers, allLocations, allVisits

}

func main() {
	listOfFiles := getListOfFiles()
	allUsers, allLocations, allVisits := getBasesData(listOfFiles)
	//fmt.Println(allUsers)
	//fmt.Println(allLocations)
	//fmt.Println(allVisits)
}
