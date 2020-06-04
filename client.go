package main

import (

	"os"
	"./API"

)

func getpowerlevel(string name) {
	
	API.aven(name)
	API.ant(name)
	API.mut(name)
	
}

func main() {
	var name1 string
	argsprog := os.Args
	name1 = os.Args[1]
	getpowerlevel(name)
}	
