//GoLang Program to build an Mutants API


package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "math/rand"
    "time"
)
type Mutants struct{
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
type mutant []Mutants

func allMutants(w http.ResponseWriter, r *http.Request) {
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
    	a := mutant{
        Mutants{"Apocalypse",value1},Mutants{"Professor X",value2},Mutants{"Dark Pheonix",value3},Mutants{"Magneto",value4},Mutants{"Wolverin",value5},Mutants{"Mystique",value6},
	}

    fmt.Println("Endpoint Hit:All Mutants Endpoint")
    json.NewEncoder(w).Encode(a)
  
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Homepage Endpoint Hit")
}

func handleRequests() {
    http.HandleFunc("/",homePage)
    http.HandleFunc("/mutant",allMutants)
    log.Fatal(http.ListenAndServe(":8082",nil))
}

func main() {
   handleRequests()
}
//Output:-Execute the go program and go to browser and type localhost:8082/mutant to get output and again refresh the page after 10 sec to get new value of the power.