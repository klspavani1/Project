package API

import (

	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"math/rand"
)

type Avengers struct {
	Nam string `json:"Avengers"`
	Avengers []Ch `json:"Character"`
}

type Ch struct {
	Name string `json:"name"`
	Power int `json:"max_power"`

}


func aven(name string) {
	var value1 int
	var p int
	t:=time.Second
	//Tick is used to execute the program for every specified time interval(10 sec in this program)
	s :=time.Tick(10*t)
	for range s{
		file,_:=ioutil.ReadFile("Avengers.json")
		data:=Avengers{}
		_=json.Unmarshal([]byte(file),&data)

		for i:=0;i<len(data.Avengers);i++ {
			n :=data.Avengers[i].Name
			rand.Seed(time.Now().UnixNano())
			value1=rand.Intn(100)

			if (name==n) {

				p = value1
			}
		break	
			
	}
	break
	fmt.Println(p)

}
}	





