package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strings"
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
    // testVerifTelephone()
    // testVerifEmail()
    //testLivres()
    //testTableauxDynamiques()
    //testHelloWorldArg()
    //testHelloWorldPrompt()
    //testBonjourLangue()
    testBonjourDateHeure()
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

// --- 5. Regex ---
func verifierTelephone(telephone string) bool {
    regex := `^(0|\+33)([ .-]?[1-9])([ .-]?[0-9]){8}$`
    match, _ := regexp.MatchString(regex, telephone)
    return match
}

func verifierEmail(email string) bool {
    regex := `^[a-zA-Z0-9]+(?:\.[a-zA-Z0-9]+)*@[a-zA-Z0-9]+\.(fr|org|com)$`
    match, _ := regexp.MatchString(regex, email)
    return match
}

func testVerifTelephone() {
    numeros := []string{
        "0123456789",         // OK
        "01 23 45 67 89",     // OK
        "01.23.45.67.89",     // OK
        "+33 1 23 45 67 89",  // OK
        "01--23--45--67--89", // KO
        "01..23..45..67..89", // KO
        "012345",             // KO
        "+33123456789",       // OK
    }

    fmt.Println("--- Test Téléphones ---")
    for _, num := range numeros {
        fmt.Printf("%-20s -> %v\n", num, verifierTelephone(num))
    }
}

func testVerifEmail() {
    emails := []string{
        "prenom.nom@email.fr",   // OK
        "user123@site.com",      // OK
        "utilisateur@domaine.org", // OK
        ".user@site.com",        // KO
        "user.@site.com",        // KO
        "user..name@site.fr",    // KO
        "user@site.net",         // KO
        "user@site",             // KO
        "user@.com",             // KO
    }

    fmt.Println("--- Test Emails ---")
    for _, email := range emails {
        fmt.Printf("%-30s -> %v\n", email, verifierEmail(email))
    }
}


// --- 6. Les livres ---
type Livre struct {
    ID          int
    Titre       string
    Auteur      string
    Description string
}

func NouveauLivre(id int, titre, auteur, description string) Livre {
    return Livre{
        ID:          id,
        Titre:       titre,
        Auteur:      auteur,
        Description: description,
    }
}

func AfficherDetails(l Livre) {
    fmt.Printf("ID: %d\nTitre: %s\nAuteur: %s\nDescription: %s\n\n", l.ID, l.Titre, l.Auteur, l.Description)
}

type Bibliotheque struct {
    Livres []Livre
}

func (b *Bibliotheque) AjouterLivre(l Livre) {
    b.Livres = append(b.Livres, l)
}

func (b Bibliotheque) AfficherLivres() {
    for _, livre := range b.Livres {
        AfficherDetails(livre)
    }
}

func (b Bibliotheque) RechercherParID(id int) *Livre {
    for _, livre := range b.Livres {
        if livre.ID == id {
            return &livre
        }
    }
    return nil
}

func testLivres() {
    fmt.Println("--- Test Livres ---")
    livre1 := NouveauLivre(1, "1984", "George Orwell", "Dystopie sur une société totalitaire.")
    livre2 := NouveauLivre(2, "Le Petit Prince", "Antoine de Saint-Exupéry", "Conte philosophique.")

    biblio := Bibliotheque{}
    biblio.AjouterLivre(livre1)
    biblio.AjouterLivre(livre2)

    fmt.Println("Tous les livres dans la bibliothèque :")
    biblio.AfficherLivres()

    fmt.Println("Recherche du livre avec ID 2 :")
    livreTrouve := biblio.RechercherParID(2)
    if livreTrouve != nil {
        AfficherDetails(*livreTrouve)
    } else {
        fmt.Println("Livre non trouvé.")
    }
}

// --- 7. Interface du tableau dynamique ---
type TableauDynamique interface {
    Ajouter(valeur interface{})
    Obtenir(index int) interface{}
}

type TableauDoublement struct {
    donnees []interface{}
    taille  int
    capacite int
}

func NouveauTableauDoublement() *TableauDoublement {
    return &TableauDoublement{
        donnees: make([]interface{}, 0, 1),
        taille:  0,
        capacite: 1,
    }
}

func (t *TableauDoublement) Ajouter(valeur interface{}) {
    if t.taille >= t.capacite {
        t.capacite *= 2
        nouveau := make([]interface{}, t.taille, t.capacite)
        copy(nouveau, t.donnees)
        t.donnees = nouveau
    }
    t.donnees = append(t.donnees, valeur)
    t.taille++
}

func (t *TableauDoublement) Obtenir(index int) interface{} {
    if index >= 0 && index < t.taille {
        return t.donnees[index]
    }
    return nil
}

type TableauAgrandissementUnitaire struct {
    donnees []interface{}
    taille  int
}

func NouveauTableauAgrandissementUnitaire() *TableauAgrandissementUnitaire {
    return &TableauAgrandissementUnitaire{
        donnees: make([]interface{}, 0),
        taille:  0,
    }
}

func (t *TableauAgrandissementUnitaire) Ajouter(valeur interface{}) {
    nouveau := make([]interface{}, t.taille+1)
    copy(nouveau, t.donnees)
    nouveau[t.taille] = valeur
    t.donnees = nouveau
    t.taille++
}

func (t *TableauAgrandissementUnitaire) Obtenir(index int) interface{} {
    if index >= 0 && index < t.taille {
        return t.donnees[index]
    }
    return nil
}

func testTableauxDynamiques() {
    fmt.Println("--- Test TableauDoublement ---")
    tab1 := NouveauTableauDoublement()
    for i := 0; i < 10; i++ {
        tab1.Ajouter(i * 10)
        fmt.Printf("Ajouté %d, taille: %d, capacité: %d\n", i*10, tab1.taille, tab1.capacite)
    }
    fmt.Println("Valeurs dans TableauDoublement :")
    for i := 0; i < tab1.taille; i++ {
        fmt.Println(tab1.Obtenir(i))
    }

    fmt.Println("\n--- Test TableauAgrandissementUnitaire ---")
    tab2 := NouveauTableauAgrandissementUnitaire()
    for i := 0; i < 10; i++ {
        tab2.Ajouter(i * 100)
        fmt.Printf("Ajouté %d, taille: %d\n", i*100, tab2.taille)
    }
    fmt.Println("Valeurs dans TableauAgrandissementUnitaire :")
    for i := 0; i < tab2.taille; i++ {
        fmt.Println(tab2.Obtenir(i))
    }
}

// --- 8. Nuances de “Hello world !” ---
func helloWorld(){
    fmt.Println("Hello world !")
}

func testHelloWorldArg() {
    nom := flag.String("nom", "inconnu", "Nom de l'utilisateur")
    flag.Parse()
    fmt.Printf("Hello %s !\n", *nom)
}

func testHelloWorldPrompt() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Entrez votre nom : ")
    nom, _ := reader.ReadString('\n')
    nom = strings.TrimSpace(nom)
    fmt.Printf("Bonjour %s !\n", nom)
}


func MessageParLangue(code string) string {
    if code == "" {
        return "Code langue vide."
    }

    messages := map[string]string{
        "fr": "Bonjour !",
        "en": "Hello!",
        "es": "¡Hola!",
        "de": "Hallo!",
        "it": "Ciao!",
        "pt": "Olá!",
        "ru": "Привет!",
        "zh": "你好！",
        "ja": "こんにちは！",
        "kv": "Видза оланныд!",
        "ar": "مرحبا!",
    }

    if msg, ok := messages[code]; ok {
        return msg
    }
    return fmt.Sprintf("Code langue inconnu : '%s'", code)
}

func testBonjourLangue() {
    code := flag.String("lang", "kv", "Code langue (ex: fr, en, es, de, it)")
    flag.Parse()
    fmt.Println(MessageParLangue(*code))
}

func MessageSelonHeure() string {
    now := time.Now()
    heure := now.Hour()
    var moment string
    switch {
    case heure >= 5 && heure < 12:
        moment = "matin"
    case heure >= 12 && heure < 17:
        moment = "après-midi"
    case heure >= 17 && heure < 21:
        moment = "soir"
    default:
        moment = "nuit"
    }
    return fmt.Sprintf("Nous sommes le %s. Bonne %s !", now.Format("lundi 2 janvier 2006"), moment)
}

func testBonjourDateHeure() {
    fmt.Println(MessageSelonHeure())
}
