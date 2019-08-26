package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("TEST")
	file.Close()

	stream, err := ioutil.ReadFile("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	aux := string(stream)
	fmt.Println(aux)

}
