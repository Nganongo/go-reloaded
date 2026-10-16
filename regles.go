package main

import (
	"strconv"
	"strings"
	"unicode"
)

// tranformer applique le marquer sur chaque ligne
func transformer(texte string) string {
	lignes := strings.Split(texte, "\n")
	for i, ligne := range lignes {
		ligne = strings.TrimRight(ligne, "\r")
		mots := strings.Fields(ligne)
		mots = appliquerMarqueur(mots)
		mots = appliquerPonctuation(mots)
		lignes[i] = strings.Join(mots, " ")
	}
	return strings.Join(lignes, "\n")
}

func appliquerMarqueur(mots []string) []string {
	var result []string
	for i := 0; i < len(mots); i++ {
		mot := mots[i]
		if (mot == "(up," || mot == "(low," || mot == "(cap,") && i+1 < len(mots) {
			suivant := mots[i+1]
			texte := strings.TrimSuffix(suivant, ")")
			n, err := strconv.Atoi(texte)
			if err == nil {
				debut := len(result) - n
				if debut < 0 {
					debut = 0
				}
				for j := debut; j < len(result); j++ {
					if mot == "(up," {
						result[j] = strings.ToUpper(result[j])
					} else if mot == "(low," {
						result[j] = strings.ToLower(result[j])
					} else if mot == "(cap," {
						result[j] = capitaliser(result[j])
					}
				}
				i++
				continue
			}
		}
		if mot == "(up)" || mot == "(low)" || mot == "(cap)" || mot == "(hex)" || mot == "(bin)" {
			if len(result) == 0 {
				continue
			}
		}
		dernier := len(result) - 1
		if mot == "(up)" {
			result[dernier] = strings.ToUpper(result[dernier])
		} else if mot == "(low)" {
			result[dernier] = strings.ToLower(result[dernier])
		} else if mot == "(cap)" {
			result[dernier] = capitaliser(result[dernier])
		} else if mot == "(hex)" {
			result[dernier] = hexTodec(result[dernier])
		} else if mot == "(bin)" {
			result[dernier] = binTodec(result[dernier])
		} else {
			result = append(result, mot)
		}

	}
	return result

}

// capitaliser change les premiere lettre en majuscule
func capitaliser(mot string) string {
	if mot == "" {
		return mot
	}
	lettre := []rune(strings.ToLower(mot))
	lettre[0] = unicode.ToUpper(lettre[0])
	return string(lettre)
}

// hexTodec convertit l'hexa en decimal
func hexTodec(mot string) string {
	valeur, err := strconv.ParseInt(mot, 16, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(valeur, 10)
}

// binTodec converti le binaire en decimal
func binTodec(mot string) string {
	valeur, err := strconv.ParseInt(mot, 2, 64)
	if err != nil {
		return mot
	}
	return strconv.FormatInt(valeur, 10)
}

func ponctuation(r rune) bool {
	return r == '.' || r == ',' || r == '!' || r == '?'
}
func appliquerPonctuation(mots []string) []string {
	var result []string
	for _, mot := range mots {
		result = append(result, mot)
	}
	return result
}
