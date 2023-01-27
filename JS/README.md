## JS

* Pour ce projet, j'ai été amené à réaliser une **interface en ligne de commande interactive**.
* J'ai réalisé ce projet par l'intermédiaire de **Node.js**, un shell en javascript.
* Le projet est assez ouvert, nottament sur le contenu du programme, j'ai choisi une interface qui permet la conversion d'élément de base 2 vers base 10 et inversement.
* Ce programme se trouve dans le fichier `app.js`
* Executer la ligne suivante dans le shell : `node app.js`
* Liste des éléments contituant mon programme : 
  * les fonctions `de10a2` et `de2a10` permettent de convertir un nombre de la base 10 (décimale) vers la base 2 (binaire) et inversement. 
  * La bibliothèque **readline** pour permettre à l'utilisateur de saisir des commandes et des nombres via la console.
  * La méthode **rl.prompt()**, elle permet à l'utilisateur de saisir son choix de conversion (*base 10 vers base 2 ou inversement*)
  * La fonciton **rl.close()** est appelée pour fermer l'interface de lecture en ligne
  * La fonction **process.exit()** est appelée pour quitter le programme.
