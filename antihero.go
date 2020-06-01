//GoLang Program to build an AntiHeros API

package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
)
type AntiHeroes struct{
    Name string `json:"Name"`
    Power int `json:"Power"`
    
}
//Specified values from 1 to 6 are used to store random values for every 10 sec interval
var value1 int
var value2 int
var value3 int
var value4 int
var value5 int
var value6 int
type hero []AntiHeroes

func allAntiHeroes(w http.ResponseWriter, r *http.Request) {
	t:=time.Second
	//Tick is used to execute the program for every specified time interval(10 sec in this program)
	s :=time.Tick(10*t)
	for range s{
	value1 =rand.Intn(100)
	value2 =rand.Intn(100)
	value3 =rand.Intn(100)
	value4 =rand.Intn(100)
	value5 =rand.Intn(100)
	value6 =rand.Intn(100)
	break
	}
    	a := hero{
        AntiHeroes{"Mandrin",value1},AntiHeroes{"Thanos",value2},AntiHeroes{"Galactus",value3},AntiHeroes{"Hela",value4},AntiHeroes{"Ego",value5},AntiHeroes{"Dr.Doom",value6},
	}

    fmt.Println("Endpoint Hit:All AntiHeroes Endpoint")
    json.NewEncoder(w).Encode(a)
  
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
    http.HandleFunc("/",homePage)
    http.HandleFunc("/hero",allAntiHeroes)
    log.Fatal(http.ListenAndServe(":8083",nil))
}

func main() {
   handleRequests()
}


//Output:-Execute the go program and go to browser and type localhost:8083/hero to get output and again refresh the page after 10 sec to get new value of the power.