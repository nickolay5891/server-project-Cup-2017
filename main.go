package main

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
}
