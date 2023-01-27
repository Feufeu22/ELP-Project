## GO
* J'ai un programme qui calcul un produit de matrice, avec des **goroutines** pour obtenir une exécution plus légère.
* Tous le code est dans le fichier **prodMatrices.go**.
* Dans un premier temps, j'ai du définir la structure d'une matrice
* Puis j'ai implémenté les différentes fonctions nécessaire pour le projet :
  * **Multiply** : Retourne la multiplication des deux matrices donées en paramètre
  * **Worker** : Effectue des multiplication en parallèle avec des goroutines
  * **RandomMatrix** : Retourne une matrice à données aléatoires *( dans mon exemple je prends des matrices carrées d'ordre 3 )*
* Enfin dans la fonction main, je fais les différentes opérations necessaires pour effectuer les goroutines
  * Création des *channels* & de *workers*
  * Envoie des matrices au channel jobs
* Je peux modifier le nombre de channel jobs ainsi que le nombre de worker, par exemple avec chacun de ces paramètre à **10**, j'ai le résultat suivant losque j'exécute mon programme : 

![Alt text](https://raw.githubusercontent.com/tfeutren/ELP-Project/master/GO/terminal.png)


