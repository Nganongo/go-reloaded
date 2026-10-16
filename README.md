# go-reloaded

Outil de correction et de modification de texte écrit en Go.
Le programme lit un fichier texte, applique une série de règles de transformation, puis écrit le résultat dans un nouveau fichier.

## Utilisation

```bash
go run . <fichier_entree> <fichier_sortie>
```

Exemple :

```bash
go run . test/sample.txt test/resultat.txt
```

- Le programme attend **exactement deux arguments**.
- Le fichier d'entrée et le fichier de sortie doivent être **différents**.

## Les règles

### Conversions de nombres

| Marqueur | Effet | Exemple |
|---|---|---|
| `(hex)` | Convertit le mot précédent (hexadécimal) en décimal | `1E (hex) files` → `30 files` |
| `(bin)` | Convertit le mot précédent (binaire) en décimal | `10 (bin) years` → `2 years` |

### Majuscules et minuscules

| Marqueur | Effet | Exemple |
|---|---|---|
| `(up)` | Met le mot précédent en majuscules | `go (up)` → `GO` |
| `(low)` | Met le mot précédent en minuscules | `SHOUTING (low)` → `shouting` |
| `(cap)` | Met une majuscule à la première lettre du mot précédent | `bridge (cap)` → `Bridge` |

Avec un nombre, la règle s'applique aux **n mots précédents** :

```
This is so exciting (up, 2)        → This is SO EXCITING
IT WAS THE AGE OF FOOLISHNESS (low, 3) → IT WAS THE age of foolishness
welcome to the brooklyn bridge (cap, 3) → welcome to The Brooklyn Bridge
```

Si le nombre est plus grand que le nombre de mots disponibles, tous les mots précédents sont modifiés.

### Ponctuation

Les signes `.` `,` `!` `?` `:` `;` sont collés au mot précédent et suivis d'un espace.
Les groupes comme `...` ou `!?` restent ensemble.

```
I was sitting over there ,and then BAMM !!  → I was sitting over there, and then BAMM!!
Punctuation tests are ... kinda boring ,what do you think ?
→ Punctuation tests are... kinda boring, what do you think?
```

### Apostrophes

Les apostrophes `'` vont par paire et sont collées aux mots qu'elles entourent.
Les apostrophes à l'intérieur d'un mot (`don't`) ne sont pas modifiées.

```
I am exactly how they describe me: ' awesome '
→ I am exactly how they describe me: 'awesome'
```

### Article `a` → `an`

`a` devient `an` (et `A` devient `An`) quand le mot suivant commence par une voyelle (`a`, `e`, `i`, `o`, `u`) ou par un `h`.

```
There it was. A amazing rock!  → There it was. An amazing rock!
```

> Cette règle est prévue pour des textes en anglais. En français, le verbe « a » (`elle a eu`) serait aussi transformé.

## Ordre d'application

Pour chaque ligne du fichier, les règles sont appliquées dans cet ordre :

1. **Marqueurs** : `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` et leurs versions avec nombre
2. **Apostrophes**
3. **Ponctuation**
4. **Article** `a` → `an`

Les marqueurs passent en premier parce que `(up, 2)` contient une virgule qui ne doit pas être traitée comme de la ponctuation.
Les apostrophes passent avant la ponctuation pour que `' texte ' .` donne bien `'texte'.`.

## Structure du projet

```
go-reloaded/
├── main.go       # Point d'entrée : vérifie les arguments, lit, transforme, écrit
├── fichiers.go   # lecture() et ecriture() des fichiers
├── regles.go     # Toutes les règles de transformation
├── go.mod
└── test/         # Fichiers de test
```

### Les fonctions principales (`regles.go`)

| Fonction | Rôle |
|---|---|
| `transformer` | Découpe le texte en lignes puis en mots, et applique toutes les règles |
| `appliquerMarqueur` | Gère `(hex)`, `(bin)`, `(up)`, `(low)`, `(cap)` et `(up, n)`… |
| `appliqueApostrophe` | Colle les apostrophes aux mots qu'elles entourent |
| `appliquerPonctuation` | Colle la ponctuation au mot précédent |
| `appliquerArticle` | Transforme `a` en `an` devant une voyelle ou un `h` |
| `capitaliser` | Met la première lettre en majuscule et le reste en minuscules |
| `hexTodec`, `binTodec` | Conversions vers le décimal |
| `ponctuation`, `estPonctuation` | Reconnaissent un caractère / un mot de ponctuation |
| `estVoyelle` | Reconnaît une voyelle ou un `h` |

## Tester

Exemple de fichier d'entrée :

```
it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief
Simply add 42 (hex) and 10 (bin) and you will see the result is 68.
There is no greater agony than bearing a untold story inside you.
Punctuation tests are ... kinda boring ,what do you think ?
```

Résultat attendu :

```
It was the best of times, it was the worst of TIMES, it was the age of wisdom, It Was The Age Of Foolishness, it was the epoch of belief
Simply add 66 and 2 and you will see the result is 68.
There is no greater agony than bearing an untold story inside you.
Punctuation tests are... kinda boring, what do you think?
```

Commande pour comparer :

```bash
go run . test/sample.txt test/resultat.txt && cat test/resultat.txt
```

## Auteur

Jordan Nganongo
