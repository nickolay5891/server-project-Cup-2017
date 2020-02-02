package main

import "fmt"

func main() {
	listOfFiles := getListOfFiles()
	//allUser, allLocatoin, allVisit := getBasesData(listOfFiles)
	allUser, allLocations, allVisits := getBasesData(listOfFiles)
	fmt.Println(allUser.getUser(1))
	fmt.Println(allUser.getUser(55000))
	fmt.Println(allLocations.getLocation(1))
	fmt.Println(allLocations.getLocation(55000))
	fmt.Println(allVisits.getVisit(1))
	fmt.Println(allVisits.getVisit(700000))
}
