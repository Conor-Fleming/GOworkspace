package main

import (
	"encoding/json"
	"fmt"
)

//taking in and "unmarshaling" json format w go
func unmarshal() {
	s := `[{"First":"James","Last":"Bond","Age":32},{"First":"Ciara","Last":"Kyne","Age":23}]`
	bs := []byte(s)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	people := []person{}
	err := json.Unmarshal(bs, &people) //"pointed to" using the address @ where the data of people is stored
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("all the data", people)

}
