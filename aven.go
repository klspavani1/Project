//GoLang Program to build an Avengers API

package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
)
type Avengers struct{
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
type aven []Avengers
func allAvengers(w http.ResponseWriter, r *http.Request) {
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
    	a := aven{
        Avengers{"Iron man",value1},Avengers{"Captain America",value2},Avengers{"Spider man",value3},Avengers{"Black Panther",value4},Avengers{"Vision",value5},Avengers{"Hawk eye",value6},
}
    fmt.Println("Endpoint Hit:All Avengers Endpoint")
    json.NewEncoder(w).Encode(a)
    
  
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func main() {
  http.HandleFunc("/",homePage)
    http.HandleFunc("/aven",allAvengers)
    log.Fatal(http.ListenAndServe(":8081",nil))
}



//Output:-Execute the go program and go to browser and type localhost:8081/aven to get output and again refresh the page after 10 sec to get new value of the power.