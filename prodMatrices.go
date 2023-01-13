package main

import "fmt"

// Création de la structure Matrice :
type Matrix struct {
	Rows int
	Columns int
	Data [][]int
}

// Création de la fonction qui multiplie deux matrices
func Multiply(matrix1 Matrix, matrix2 Matrix) (Matrix, error) {

	// Vérifier si le produit est possible 
    if matrix1.Columns != matrix2.Rows {
		return Matrix{}, fmt.Errorf("Les matrices ne sont pas multipliable entre eux")
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
	return result,nil

}

func main() {

	// Créer une matrice
	A := Matrix{Rows: 2, Columns: 2, Data: [][]int{{1, 2}, {3,4}}}
	B := Matrix{Rows: 2, Columns: 2, Data: [][]int{{3,1}, {-1,2}}}

	// Afficher AxB
    fmt.Println(Multiply(A, B))
}
