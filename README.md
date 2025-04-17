# go_td1

## Exercice 4.3

### 1. Quelle est la différence entre une liste chaînée simple et une liste chaînée double ?

- Une liste chaînée simple est une structure de données où chaque nœud contient un pointeur vers le nœud suivant.
- Une liste chaînée double est une structure de données où chaque nœud contient deux pointeurs : un vers le nœud suivant et un vers le nœud précédent. Cela permet une navigation bidirectionnelle.

### 2. Que se passe-t-il si vous essayez de supprimer un élément qui n'existe pas dans la liste ?

Dans ce cas, la fonction Supprimer parcourra la liste jusqu'à la fin sans trouver l'élément et ne fera aucune modification à la liste.

### 3. Quelle est la complexité en temps des différentes opérations (ajout, suppression, recherche) sur une liste chaînée ?

- Ajout : O(n) dans le pire des cas, car il faut parcourir la liste jusqu'à la fin pour ajouter un nouvel élément.
- Suppression : O(n) dans le pire des cas, car il faut parcourir la liste pour trouver l'élément à supprimer.
- Recherche : O(n) dans le pire des cas, car il faut parcourir la liste pour trouver l'élément recherché.

Réalisé par Margarita Surina