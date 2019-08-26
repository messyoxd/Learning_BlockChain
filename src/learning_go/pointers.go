package main

import "fmt"

func main() {

	x := "aaaaa"

	mudaX(&x)

	fmt.Println(x)
}

func mudaX(x *string) {
	*x = "bbbbbb"
}
