package API

import (
"encoding/json"
"fmt"
"io/ioutil"
"time"
"math/rand"
)

type Anti struct {
	Nam string `json:"Anti"`
	Anti []An `json:"Character"`
}

type An struct {
	Name string `json:"name"`
	Power int `json:"max_power"`
}

func ant(name string) {
	var value2 int 
	var p1 int 
	t1:=time.Second
	//Tick is used to execute the program for every specified time interval(10 sec in this program)
	s1 :=time.Tick(10*t1)
	for range s1{
		file,_:=ioutil.ReadFile("Anti.json")
		data:=Anti{}
		_=json.Unmarshal([]byte(file),&data)

		for i:=0;i<len(data.Anti);i++ {

			n1 :=data.Anti[i].Name
			rand.Seed(time.Now().UnixNano())
			value2=rand.Intn(100)

			if (name==n1) {

				p1 = value2
			}
		break	
			
	}
	break
	fmt.Println(p1)

}
}
	





