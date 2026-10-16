package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("veillez entrer deux fichier ")
		os.Exit(1)
	}

	entree := os.Args[1]
	sortie := os.Args[2]
	if entree == sortie {
		fmt.Fprintln(os.Stderr, "Erreur : le fichier d'entrée et le fichier de sortie doivent être différents")
		os.Exit(1)
	}

	contenue, err := lecture(entree)
	if err != nil {
		fmt.Println("Erreur", err)
		os.Exit(1)
	}
	resultat := transformer(contenue)
	err = ecriture(sortie, resultat)
	if err != nil {
		fmt.Println("erreur", err)
		os.Exit(1)
	}

	fmt.Println("Terminé:", sortie, "a été creé")

}
