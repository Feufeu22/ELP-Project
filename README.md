# PROJET ELP

* **Objectif** : se familiariser avec certains de ces concepts en utilisant différents langages de programmation
* Le projet est constitué de trois parties indépendantes ( *ELM, GO , JS* )

## ELM

* L'objectif de ce projet est d'écrire un programme elm d'une application web qui permet de deviner un mot dont les définitions sont affichées.
* Pour ce faire, nous utiliser l'API :  [Free Dictionary API](https://dictionaryapi.dev/).  
* Puis nous utiliser la base de données suivante es mots se trouve dans ce [fichier](https://perso.liris.cnrs.fr/tristan.roussillon/GuessIt/thousand_words_things_explainer.txt) qui correspond à 1000 mots dans un fichier texte.

### Ouverture du projet en local :
* Se diriger vers le fichier ELM/src
* Taper la commande suivante : 
```bash
elm reactor
```
* Naviguer en serveur local à partir de l'adresse suivante : http://localhost:8000
* Se diriger vers le fichier `Projet.elm`.

## GO

* Implantation d'une interaction client-serveur pour le calcul de matrice
* ``` server.go ``` pour le serveur et ``` client.go ``` pour le client
* Côté serveur :
```bash
go run server.go <port>
```
* Côté client :
```bash
go run client.go <port> <input>
```
* La sortie est écrite dans ``` output.txt ```
* Utilisation d'une goroutine qui effectue des tâches en continu en récupérant des travaux à partir d'un canal
* Calculs matriciels distribués en utilisant des workers et traite les demandes provenant des client.

## JS

* CLI en javascript
* accèder au fichier : /JS/bashJS/bin/bash.js
* Commande à exécuter :
```bash
bash bash.js
```
* Les commandes accepter pour le CLI : 
  * exit
  * clear
  * help
  * ls
  * cd 
  * touch
  * mkdir
  * rm


