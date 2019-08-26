package main

import "fmt"

func main() {
	var nome string = "ronaldo"
	const nome2 string = "aaaaa"
	fmt.Println(nome, " ", nome2)

	var a int = 30
	var b float64 = 30.0003
	c, d := 40.0001, "bbbbbb"
	fmt.Println(a, ",", b, ",", c, ",", d)

	is_false := false
	fmt.Println(is_false)

	fmt.Printf("%d %f %t %s \n", a, b, is_false, d)

	fmt.Println(add(b, c))
}

func add(a float64, b float64) float64 {
	return a + b
}
