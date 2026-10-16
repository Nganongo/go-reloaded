package main

import (
	"fmt"
	"os"
)

// lecture: pour lire le fichier entréé
func lecture(chemin string) (string, error) {
	contenue, err := os.ReadFile(chemin)
	if err != nil {
		return "", fmt.Errorf("impossible de lire le fichier %s", chemin)
	}
	return string(contenue), nil

}

func ecriture(chemin, texte string) error {
	err := os.WriteFile(chemin, []byte(texte), 0644)
	if err != nil {
		return fmt.Errorf("impossible d'ecrire sur le fichier %s", chemin)
	}
	return nil

}
