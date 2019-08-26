package main

import "fmt"

func main() {
	pessoaMap := make(map[string]int)

	pessoaMap["eu"] = 1
	pessoaMap["tu"] = 2

	fmt.Println(pessoaMap)

	fmt.Println(pessoaMap["eu"])

	fmt.Println(len(pessoaMap))

	/* map inside map */
	pessoa2Map := map[string]map[string]string{
		"eu": map[string]string{
			"idade": "21",
		},
	}
	if t, pessoa := pessoa2Map["eu"]; pessoa {
		fmt.Println(t["idade"])
	}
}
