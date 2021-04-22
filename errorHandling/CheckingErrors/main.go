package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	n, err := fmt.Println("Hello")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(n)

	f, err := os.Create("error.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	r := strings.NewReader("Random Errors not happening")
	io.Copy(f, r)

	x, err := os.Open("error.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer x.Close()

	bs, err := ioutil.RealAll(x)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
