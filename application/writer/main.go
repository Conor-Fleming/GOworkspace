package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("hello")
	fmt.Fprintln(os.Stdout, "hello")
	io.WriteString(os.Stdout, "hello")

}
