package main

import (
	"fmt"
	"log"
	"os"
)

//error checking with printing errors and writing to log files
func main() {
	f, err2 := os.Create("log.txt")
	if err2 != nil {
		fmt.Println(err2)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("error occured, no such file when --->", err) //print message to command line
		log.Println("error occured, no such file when --->", err) //print err to a log file created above
		//log.Fatalln(err) //calls os.Exit(1)...program will end
		log.Panicln(err) //prints panic message with exit status of 2
		//panic(err)
	}
	defer f2.Close()

	fmt.Println("check the Log file for a full description of error")
}
