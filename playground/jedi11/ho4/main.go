package main

import (
	"errors"
	"fmt"
	"log"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	result, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(result)
	}
}

func sqrt(f float64) (float64, error) { //return type error that looks to the struct and func above
	if f < 0 {
		return 0, sqrtError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  errors.New("this is the error ive created"), // could also use 'fmt.Errorf("**insert message here** %v", f)
		}
	}
	return f, nil
}
