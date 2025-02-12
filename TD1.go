package main

import (
    "fmt"
	"math"
    "math/rand"
    // "os"
    // "os/exec"
    "time"
)

func main() {
	// noms := []string{"Alice", "Bob", "Charlie", "David", "Eve", "Frank"}
    // groupes := organiserParTaille(noms)
    // for longueur, noms := range groupes {
    //     fmt.Printf("Longueur %d: %v\n", longueur, noms)
    // }

    // grille := initGrille(5, 5)
    // for _, ligne := range grille {
    //     fmt.Println(ligne)
    // }

	// fmt.Println("Nombre de voisins vivants autour de la cellule :", compterVoisins(grille, 0, 0))

	// grille = update(grille)
    // fmt.Println("Grille mise à jour:")
    // for _, ligne := range grille {
    //     fmt.Println(ligne)
    // }

	// afficherGrille(grille)

	// //test 2.5
	// grille := initGrille(50, 50)
    // for {
    //     // Effacer l'affichage du terminal
        // c := exec.Command("clear")
        // c.Stdout = os.Stdout
        // c.Run()

    //     // Afficher la grille
    //     afficherGrille(grille)

    //     // Mettre à jour la grille
    //     grille = update(grille)

    //     // Pause de 500 millisecondes
    //     time.Sleep(500 * time.Millisecond)
    // }

    // // Initialisation de deux vecteurs 2D
    // v1 := initVec2i(3, 4)
    // v2 := initVec2i(1, 2)

    // // Addition
    // vAdd := additionVec2i(v1, v2)
    // fmt.Printf("Addition: (%d, %d)\n", vAdd.x, vAdd.y)

    // // Soustraction
    // vSub := soustractionVec2i(v1, v2)
    // fmt.Printf("Soustraction: (%d, %d)\n", vSub.x, vSub.y)

    // // Multiplication
    // vMul := multiplicationVec2i(v1, v2)
    // fmt.Printf("Multiplication: (%d, %d)\n", vMul.x, vMul.y)

    // // Norme
    // norm := normeVec2i(v1)
    // fmt.Printf("Norme: %f\n", norm)

    // // Normalisation
    // vNorm := normalisationVec2i(v1)
    // fmt.Printf("Normalisation: (%d, %d)\n", vNorm.x, vNorm.y)

    // // Produit scalaire
    // dot := produitScalaireVec2i(v1, v2)
    // fmt.Printf("Produit scalaire: %d\n", dot)

    // // Produit vectoriel
    // cross := produitVectorielVec2i(v1, v2)
    // fmt.Printf("Produit vectoriel: %d\n", cross)

    list := LinkedList{}

    // Ajouter des éléments à la liste
    list.Ajouter(1)
    list.Ajouter(2)
    list.Ajouter(3)
    list.Afficher() // Affiche: 1 -> 2 -> 3 -> nil

    // Insérer un élément à une position donnée
    list.InsertAtPosition(4, 1)
    list.Afficher() // Affiche: 1 -> 4 -> 2 -> 3 -> nil

    // Supprimer un élément de la liste
    list.Supprimer(2)
    list.Afficher() // Affiche: 1 -> 4 -> 3 -> nil

    // Supprimer un élément qui n'existe pas
    list.Supprimer(5)
    list.Afficher() // Affiche: 1 -> 4 -> 3 -> nil

    // Insérer un élément à une position au-delà de la fin de la liste
    list.InsertAtPosition(5, 10)
    list.Afficher() // Affiche: 1 -> 4 -> 3 -> 5 -> nil
}

//1
func estBissextile(annee int) bool {
	if annee % 4 == 0 {
		if annee % 100 == 0 {
			if annee % 400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func estPremier(nombre int) bool {
	for i := 2; i < nombre; i++ {
		if nombre % i == 0 {
			return false
		}
	}
	return true
}

func listeNombrePremier(n int) []int {
    var liste []int
    i := 2
    for len(liste) < n {
        if estPremier(i) {
            liste = append(liste, i)
        }
        i++
    }
    return liste
}

func genererTableauAleatoire (n int) []int{
	var liste []int
	for len(liste) < n {
		liste = append(liste, rand.Intn(100))
	}
	return liste
}

func triBulles(listeATrier []int) []int {
	if len(listeATrier) == 0 {
		return listeATrier
	} else {
		for i := 1; i < len(listeATrier); i++ {
			for j := 0; j < len(listeATrier)-i; j++ {
				if listeATrier[j+1] < listeATrier[j] {
					listeATrier[j], listeATrier[j+1] = listeATrier[j+1], listeATrier[j]
				}
			}
		}
	}
	return listeATrier
}

func triSelection(listeATrier []int) []int {
	n := len(listeATrier)
	for i:=0; i < n-2 ; i++{
		min := i
		for j:= i+1; j < n-1 ; j++ {
			if listeATrier[j] < listeATrier[min] {
				min = j
			}
		}
		if min != i {
			listeATrier[min], listeATrier[i] = listeATrier[i], listeATrier[min]
		}
	}
	return listeATrier
}

func rechercheDichotomique(liste []int, n int) (int,bool) {
	for i:= 0; i < len(liste); i++ {
		if liste[i] == n {
			return i,true
		} else {
			return -1,false
		}
			
	}
	return -1,false
}

func organiserParTaille(liste []string) map[int][]string {
    groupes := make(map[int][]string)
    for _, nom := range liste {
        longueur := len(nom)
        groupes[longueur] = append(groupes[longueur], nom)
    }
    return groupes
}

//2
// initGrille initialise une grille de taille n x m avec des cellules vivantes (1) et mortes (0) de manière aléatoire
func initGrille(n, m int) [][]int {
    rand.Seed(time.Now().UnixNano())
    grille := make([][]int, n)
    for i := range grille {
        grille[i] = make([]int, m)
        for j := range grille[i] {
            grille[i][j] = rand.Intn(2) // 0 ou 1
        }
    }
    return grille
}

func compterVoisins(grille [][]int, i, j int) int {
    n, m := len(grille), len(grille[0])
    voisins := [8][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1},         {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }
    compteur := 0
    for k := 0; k < len(voisins); k++ {
        ni, nj := i+voisins[k][0], j+voisins[k][1]
        if ni >= 0 && ni < n && nj >= 0 && nj < m && grille[ni][nj] == 1 {
            compteur++
        }
    }
    return compteur
}

func update(grille [][]int) [][]int {
    n, m := len(grille), len(grille[0])
    nouvelleGrille := make([][]int, n)
    for i := range nouvelleGrille {
        nouvelleGrille[i] = make([]int, m)
        for j := range nouvelleGrille[i] {
            voisinsVivants := compterVoisins(grille, i, j)
            if grille[i][j] == 1 && (voisinsVivants==2 || voisinsVivants == 3) {
                nouvelleGrille[i][j] = 1
            } else if grille[i][j] == 0 && voisinsVivants == 3 {
                nouvelleGrille[i][j] = 1
            } else {
                nouvelleGrille[i][j]=0
            }
        }
    }
    return nouvelleGrille
}

func afficherGrille(grille [][]int) {
    for i := range grille {
        for j := range grille[i] {

            if grille[i][j] == 0 {
                fmt.Print("\u2588")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

//3
type vec2i struct {
    x int
    y int
}

func initVec2i(x, y int) vec2i {
    return vec2i{x: x, y: y}
}

func additionVec2i(v, v1 vec2i) vec2i{
	return vec2i{x: v.x + v1.x, y: v.y + v1.y}
}

func soustractionVec2i(v, v1 vec2i) vec2i {
    return vec2i{x: v.x - v1.x, y: v.y - v1.y}
}

func multiplicationVec2i(v, v1 vec2i) vec2i {
    return vec2i{x: v.x * v1.x, y: v.y * v1.y}
}

func normeVec2i(v vec2i) float64 {
    return math.Sqrt(float64(v.x*v.x + v.y*v.y))
}

func normalisationVec2i(v vec2i) vec2i {
    norm := normeVec2i(v)
    if norm == 0 {
        return vec2i{0, 0}
    }
    return vec2i{x: int(float64(v.x) / norm), y: int(float64(v.y) / norm)}
}

func produitScalaireVec2i(v, v1 vec2i) int {
    return v.x*v1.x + v.y*v1.y
}

func produitVectorielVec2i(v, v1 vec2i) int {
    return v.x*v1.y - v.y*v1.x
}

//4
// Node représente un nœud dans une liste chaînée
type Node struct {
    Data int
    Next *Node
}

// LinkedList représente une liste chaînée
type LinkedList struct {
    Head *Node
}

// Ajouter ajoute un nouvel élément à la fin de la liste chaînée
func (list *LinkedList) Ajouter(data int) {
    newNode := &Node{Data: data}
    if list.Head == nil {
        list.Head = newNode
    } else {
        current := list.Head
        for current.Next != nil {
            current = current.Next
        }
        current.Next = newNode
    }
}

// Supprimer supprime le premier nœud contenant l'élément spécifié
func (list *LinkedList) Supprimer(data int) {
    if list.Head == nil {
        return
    }
    if list.Head.Data == data {
        list.Head = list.Head.Next
        return
    }
    current := list.Head
    for current.Next != nil && current.Next.Data != data {
        current = current.Next
    }
    if current.Next != nil {
        current.Next = current.Next.Next
    }
}

// Afficher affiche tous les éléments de la liste chaînée
func (list *LinkedList) Afficher() {
    current := list.Head
    for current != nil {
        fmt.Printf("%d -> ", current.Data)
        current = current.Next
    }
    fmt.Println("nil")
}

// InsertAtPosition insère un élément à une position donnée dans la liste chaînée
func (list *LinkedList) InsertAtPosition(data int, position int) {
    newNode := &Node{Data: data}
    if position == 0 {
        newNode.Next = list.Head
        list.Head = newNode
        return
    }
    current := list.Head
    for i := 0; i < position-1 && current != nil; i++ {
        current = current.Next
    }
    if current == nil {
        list.Ajouter(data)
    } else {
        newNode.Next = current.Next
        current.Next = newNode
    }
}

/*

1. Quelle est la différence entre une liste chaînée simple et une liste chaînée double ?

- Une liste chaînée simple (simplement chaînée) est une structure de données où chaque nœud contient un pointeur vers le nœud suivant.
- Une liste chaînée double (doublement chaînée) est une structure de données où chaque nœud contient deux pointeurs : un vers le nœud suivant et un vers le nœud précédent. Cela permet une navigation bidirectionnelle.

2. Que se passe-t-il si vous essayez de supprimer un élément qui n'existe pas dans la liste ?

Si vous essayez de supprimer un élément qui n'existe pas dans la liste, la fonction Supprimer parcourra la liste jusqu'à la fin sans trouver l'élément et ne fera aucune modification à la liste.

3. Quelle est la complexité en temps des différentes opérations (ajout, suppression, recherche) sur une liste chaînée ?

- Ajout : O(n) dans le pire des cas, car il faut parcourir la liste jusqu'à la fin pour ajouter un nouvel élément.
- Suppression : O(n) dans le pire des cas, car il faut parcourir la liste pour trouver l'élément à supprimer.
- Recherche : O(n) dans le pire des cas, car il faut parcourir la liste pour trouver l'élément recherché.

*/