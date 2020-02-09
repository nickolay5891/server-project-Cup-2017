package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var (
	allUsers     AllUsers
	allLocations AllLocations
	allVisits    AllVisits
)

func main() {
	listOfFiles := getListOfFiles()
	allUsers, allLocations, allVisits = getBasesData(listOfFiles)
	//allUsers, _, _ = getBasesData(listOfFiles)
	//fmt.Println(allUsers)
	//fmt.Println(allUsers.getUser(5))
	//fmt.Println(allUsers.getUser(55000))
	//fmt.Println(allLocations.getLocation(1))
	//fmt.Println(allLocations.getLocation(55000))
	//fmt.Println(allVisits.getVisit(1))
	//fmt.Println(allVisits.getVisit(700000))
	//f := 1295267957
	//t := 1370802458
	//c := "Египет"
	//var d uint32 = 15
	//g := "m"
	//var aF uint32 = 50
	//var aT uint32 = 80
	//fmt.Println(getUserVisits(52, &f, &t, &c, &d))
	//fmt.Println(getUserVisits(52, &f, nil, nil, nil))
	//getUserVisits(52, &f, &t, &c, &d)
	//fmt.Println(getLocationsAvg(52, &f, &t, &aF, &aT, nil))
	//fmt.Println(getAgeUser(8947583))
	//fmt.Println(getAgeUser(571183261))
	//fmt.Println(getAgeUser(571356061))
	//fmt.Println(round(tf, 5))

	//u := User{"swe", "fggf", 6567, "32", 89, "hkj"}
	//allUsers.createUser(u)
	//u1 := User{"swe", "fggf", 90, "32", 15000, "hkj"}
	//allUsers.createUser(u1)
	//fmt.Println(allUsers.getUser(u1.ID))

	//l := Location{10, "Homel", "cafe", 89, "Belarus"}
	//allLocations.createLocation(l)
	//l1 := Location{10, "Homel", "cafe", 50000, "Belarus"}
	//allLocations.createLocation(l1)
	//fmt.Println(allLocations.getLocation(l1.ID))

	//v := Visit{10, 20, 40, 89, 5}
	//allVisits.createVisit(v)
	//fmt.Println(allVisits.getVisit(v.ID))
	//v1 := Visit{10, 20, 40, 700000, 5}
	//allVisits.createVisit(v1)
	//fmt.Println(allVisits.getVisit(v1.ID))

	//us := User{"swe", "fggf", 6567, "32", 89, "hkj"}

	//qw, _ := allUsers.getUser(2)
	//fmt.Println(qw)
	//getDU(&us)

	//fmt.Println(allUsers.getUser(2))
	//fmt.Println(*getDU().FirstName)
	//fmt.Println(qw.FirstName)
	//allUsers.updateUser(2, getDU())
	//fmt.Println(allUsers.getUser(2))
}

func getDU() UserPost {
	file, err := ioutil.ReadFile("us1.json")
	if err != nil {
		log.Fatal(err)
	}
	var requestBody UserPost
	json.Unmarshal(file, &requestBody)
	return requestBody
}
