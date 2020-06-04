package API

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"math/rand"
)

type Mutants struct {
	Nam string `json:"name"`
	Mutants []Mu `json:"Character"`
}

type Mu struct {
	Name string `json:"name"`
	Power int `json:"max_power"`
}


func mut(name string) {
	var value3 int 
	var p2 int
	t:=time.Second
	//Tick is used to execute the program for every specified time interval(10 sec in this program)
	s2 :=time.Tick(10*t)
	for range s2{
	
		file,_:=ioutil.ReadFile("Mutants.json")
		data:=Mutants{}
		_=json.Unmarshal([]byte(file),&data)

		for i:=0;i<len(data.Mutants);i++ {
			n2 :=data.Mutants[i].Name
			rand.Seed(time.Now().UnixNano())
			value3=rand.Intn(100)
			if (name==n2) {

				p2 = value3
			}
		break	
			
	}
	break
	fmt.Println(p2)


}
}
	

