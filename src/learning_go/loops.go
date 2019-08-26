package main

import "fmt"

func main() {

	/* for loop */
	for i := 0; i < 10; i++ {
		if i%2 == 0 && i < 10 {
			fmt.Println(i)
		} else if i == 9 {
			fmt.Println(" ")
		} else {
			fmt.Printf("")
		}
	}

	/* while loop */
	i := 0
	for i < 10 {
		i++
		switch i {
		case 10:
			fmt.Printf("%d\n\n", i)
		default:
			fmt.Println(i)
		}
	}
}
