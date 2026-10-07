package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("il manque un fichier")
		return
	}

	test := os.Args[1]
	result := os.Args[2]
	contenu, err := os.ReadFile(test)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = os.WriteFile(result, contenu, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}
