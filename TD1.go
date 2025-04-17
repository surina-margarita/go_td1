package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

// --- Main ---

func main() {
    runDemo()
}

// --- Démos à activer ici ---

func runDemo() {
    // testListeChainee()
    // testTriEtRecherche()
    // testVecteurs()
    // testJeuDeLaVie()
}

// --- 1. Premiers pas ---

func estBissextile(annee int) bool {
    if annee%4 == 0 {
        if annee%100 == 0 {
            return annee%400 == 0
        }
        return true
    }
    return false
}

func estPremier(nombre int) bool {
    for i := 2; i < nombre; i++ {
        if nombre%i == 0 {
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

func genererTableauAleatoire(n int) []int {
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

func testTriEtRecherche() {
    liste := genererTableauAleatoire(10)
    fmt.Println("Liste aléatoire :", liste)
    liste = triBulles(liste)
    fmt.Println("Tri à bulles :", liste)
    index, trouvé := rechercheDichotomique(liste, liste[3])
    fmt.Println("Recherche :", index, trouvé)
}

// --- 2. Jeu de la vie ---

func initGrille(n, m int) [][]int {
    grille := make([][]int, n)
    for i := range grille {
        grille[i] = make([]int, m)
        for j := range grille[i] {
            grille[i][j] = rand.Intn(2)
        }
    }
    return grille
}

func compterVoisins(grille [][]int, i, j int) int {
    n, m := len(grille), len(grille[0])
    voisins := [8][2]int{
        {-1, -1}, {-1, 0}, {-1, 1},
        {0, -1}, {0, 1},
        {1, -1}, {1, 0}, {1, 1},
    }
    compteur := 0
    for _, v := range voisins {
        ni, nj := i+v[0], j+v[1]
        if ni >= 0 && ni < n && nj >= 0 && nj < m && grille[ni][nj] == 1 {
            compteur++
        }
    }
    return compteur
}

func update(grille [][]int) [][]int {
    n, m := len(grille), len(grille[0])
    nouvelle := make([][]int, n)
    for i := range nouvelle {
        nouvelle[i] = make([]int, m)
        for j := range nouvelle[i] {
            voisins := compterVoisins(grille, i, j)
            if grille[i][j] == 1 && (voisins == 2 || voisins == 3) {
                nouvelle[i][j] = 1
            } else if grille[i][j] == 0 && voisins == 3 {
                nouvelle[i][j] = 1
            }
        }
    }
    return nouvelle
}

func afficherGrille(grille [][]int) {
    for _, ligne := range grille {
        for _, val := range ligne {
            if val == 0 {
                fmt.Print("█")
            } else {
                fmt.Print(" ")
            }
        }
        fmt.Println()
    }
}

func testJeuDeLaVie() {
    grille := initGrille(10, 10)
    for i := 0; i < 5; i++ {
        afficherGrille(grille)
        fmt.Println()
        grille = update(grille)
        time.Sleep(500 * time.Millisecond)
    }
}

// --- 3. Vecteurs ---

type vec2i struct {
    x, y int
}

func initVec2i(x, y int) vec2i { return vec2i{x, y} }

func additionVec2i(v, v1 vec2i) vec2i {
    return vec2i{v.x + v1.x, v.y + v1.y}
}

func soustractionVec2i(v, v1 vec2i) vec2i {
    return vec2i{v.x - v1.x, v.y - v1.y}
}

func multiplicationVec2i(v, v1 vec2i) vec2i {
    return vec2i{v.x * v1.x, v.y * v1.y}
}

func normeVec2i(v vec2i) float64 {
    return math.Sqrt(float64(v.x*v.x + v.y*v.y))
}

func normalisationVec2i(v vec2i) vec2i {
    n := normeVec2i(v)
    if n == 0 {
        return vec2i{0, 0}
    }
    return vec2i{int(float64(v.x) / n), int(float64(v.y) / n)}
}

func produitScalaireVec2i(v, v1 vec2i) int {
    return v.x*v1.x + v.y*v1.y
}

func produitVectorielVec2i(v, v1 vec2i) int {
    return v.x*v1.y - v.y*v1.x
}

func testVecteurs() {
    v1 := initVec2i(3, 4)
    v2 := initVec2i(1, 2)
    fmt.Println("Addition:", additionVec2i(v1, v2))
    fmt.Println("Soustraction:", soustractionVec2i(v1, v2))
    fmt.Println("Multiplication:", multiplicationVec2i(v1, v2))
    fmt.Println("Norme:", normeVec2i(v1))
    fmt.Println("Normalisation:", normalisationVec2i(v1))
    fmt.Println("Dot product:", produitScalaireVec2i(v1, v2))
    fmt.Println("Cross product:", produitVectorielVec2i(v1, v2))
}

// --- 4. Liste chaînée ---

type Node struct {
    Data int
    Next *Node
}

type LinkedList struct {
    Head *Node
}

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

func (list *LinkedList) Afficher() {
    for current := list.Head; current != nil; current = current.Next {
        fmt.Printf("%d -> ", current.Data)
    }
    fmt.Println("nil")
}

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

func testListeChainee() {
    list := LinkedList{}
    list.Ajouter(1)
    list.Ajouter(2)
    list.Ajouter(3)
    list.Afficher() // Affiche: 1 -> 2 -> 3 -> nil
    list.InsertAtPosition(4, 1)
    list.Afficher() // Affiche: 1 -> 4 -> 2 -> 3 -> nil
    list.Supprimer(2)
    list.Afficher() // Affiche: 1 -> 4 -> 3 -> nil
    list.Supprimer(5)
    list.Afficher()
    list.InsertAtPosition(5, 10)
    list.Afficher() // Affiche: 1 -> 4 -> 3 -> 5 -> nil
}
