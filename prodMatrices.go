/*-----------------------------------------------*/
/*               ** PROJET GO **                 */
/*-----------------------------------------------*/

package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*-----------------------------------------------*/
// ** STRUCTURES  ** //

// La structure d'une Matrice :
type Matrix struct {
	Rows    int
	Columns int
	Data    [][]int
}

// La structure d'un résultat de multiplication de matrices
type MatrixResult struct {
	result Matrix
	err    error
}

/*-----------------------------------------------*/
// ** FONCTIONS ** //

// La fonction qui multiplie deux matrices
func Multiply(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// Vérifier si le produit est possible
	if matrix1.Columns != matrix2.Rows {
		return Matrix{}, fmt.Errorf("Les matrices ne sont pas multipliable entre elles")
	}

	// Création de la matrice résultant
	result := Matrix{Rows: matrix1.Rows, Columns: matrix2.Columns}
	result.Data = make([][]int, matrix1.Rows)
	for i := range result.Data {
		result.Data[i] = make([]int, matrix2.Columns)
	}

	// Calcul de la matrice
	for i := 0; i < matrix1.Rows; i++ {
		for j := 0; j < matrix2.Columns; j++ {
			var somme int
			for k := 0; k < matrix1.Columns; k++ {
				somme += matrix1.Data[i][k] * matrix2.Data[k][j]
			}
			result.Data[i][j] = somme
		}
	}

	// Afficher la matrice résultat
	return result, nil

}

// La fonction WORKER (effectue des multiplication en parallèle avec des goroutines)
func worker(id int, jobs <-chan Matrix, results chan<- MatrixResult) {
	for j := range jobs {
		matrix1 := j
		matrix2 := <-jobs
		result, err := Multiply(matrix1, matrix2)
		if err != nil {
			results <- MatrixResult{result: result, err: err}
		} else {
			results <- MatrixResult{result: result, err: nil}
		}
	}
}

// Fonction qui génère une matrice carré d'ordre n aléatoire
func RandomMatrix(rows, columns int) Matrix {
	rand.Seed(time.Now().UnixNano())
	data := make([][]int, rows)
	for i := range data {
		data[i] = make([]int, columns)
		for j := range data[i] {
			data[i][j] = rand.Intn(10)
		}
	}
	time.Sleep(100 * time.Millisecond)
	return Matrix{Rows: rows, Columns: columns, Data: data}
}

// Fonction qui retourne les matrices à multiplier
func getMatrix1() Matrix {
	return RandomMatrix(3, 3)
}
func getMatrix2() Matrix {
	return Matrix{Rows: 3, Columns: 3, Data: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}
}

/*-----------------------------------------------*/
// ** PRINCIPAL ** //

func main() {

	// Création des channels
	jobs := make(chan Matrix)
	results := make(chan MatrixResult)

	// Création des workers
	numWorkers := 10
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// Envoyer les matrices au channel jobs
	numTasks := 10
	for i := 0; i < numTasks; i++ {
		matrix1 := getMatrix1()
		matrix2 := getMatrix2()
		jobs <- matrix1
		jobs <- matrix2
	}
	close(jobs)
	for i := 0; i < numTasks; i++ {
		res := <-results
		if res.err != nil {
			fmt.Println(res.err)
		} else {
			fmt.Println(res.result)
		}
	}
}
