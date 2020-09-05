package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func unmarsh() { //unmarshaling a json
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)
	fmt.Println("into go value00000below")
	new := []byte(s)
	peeps := []person{}
	err := json.Unmarshal(new, &peeps)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(peeps)

	for _, v := range peeps { //using range to format print values
		fmt.Println(v.First, v.Last, v.Age)
		for _, val := range v.Sayings {
			fmt.Println(val)
		}
		fmt.Println()
	}
}
