// SERVER PART //

package main

//  --------------------------------------------------
// imports
import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

//  --------------------------------------------------
// Global variables

var host string = "localhost"
var stopWorker user = user{nil, -1, -1, nil, nil}

//  --------------------------------------------------
// Constants

const size int = 100
const nb_worker int = 10

//  --------------------------------------------------
// Structures

type result struct {
	value float64
	x     int
	y     int
}

type job struct {
	x   int
	y   int
	raw *[]float64
	col *[]float64
}

type user struct {
	connection net.Conn
	sizeMatrix int
	id         int
	matrixA    [][]float64
	matrixB    [][]float64
}

//  --------------------------------------------------
// Intermediate functions

func initSquareMatrix(n int) (matrixA [][]float64, matrixB [][]float64, matrixC [][]float64) {
	A := make([][]float64, n)
	B := make([][]float64, n)
	C := make([][]float64, n)
	for i := 0; i < n; i++ {
		A[i] = make([]float64, n)
		B[i] = make([]float64, n)
		C[i] = make([]float64, n)
	}
	return A, B, C
}

func computeCoef(raw *[]float64, col *[]float64) float64 {
	var res float64
	for i := 0; i < len(*raw); i++ {
		res += (*raw)[i] * (*col)[i]
	}
	return res
}

func inputTextMatrix(text string) (matA [][]float64, matB [][]float64, matC [][]float64) {
	mat := strings.Split(text, "-")
	mattA := strings.Split(mat[0], "\n")
	mattB := strings.Split(mat[1], "\n")[1:]
	size := len(mattB)
	A, B, C := initSquareMatrix(size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			A[i][j], _ = strconv.ParseFloat(strings.Split(mattA[i], " ")[j], 3)
			B[i][j], _ = strconv.ParseFloat(strings.Split(mattB[i], " ")[j], 3)
		}
	}
	return A, B, C
}

func matrixToString(matrix [][]float64) string {
	result := ""
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			result += strconv.FormatFloat(matrix[i][j], 'f', 1, 64) + " "
		}
		result += "\n"
	}
	return result
}

func goWorker(jobChan chan job, reslutChan chan result) {
	for {
		job := <-jobChan
		if job.x == -1 && job.y == -1 {
			break
		}
		var result result
		result.x = job.x
		result.y = job.y
		result.value = computeCoef(job.raw, job.col)
		reslutChan <- result
	}
}

func goHandlerUser(newUser user) {
	defer newUser.connection.Close()
	resultChan := make(chan result, size)
	jobChan := make(chan job, size)

	for i := 0; i < nb_worker; i++ {
		go goWorker(jobChan, resultChan)
	}
	fmt.Println("Connection with a new client user")

	reader := bufio.NewReader(newUser.connection)
	data, err := reader.ReadString('$')

	if err != nil {
		fmt.Println("Error while reading", err.Error())
		return
	}

	data = strings.TrimSuffix(data, "$")
	var C [][]float64

	// Produc matrix computation
	newUser.matrixA, newUser.matrixB, C = inputTextMatrix(data)
	begin := time.Now()
	newUser.sizeMatrix = len(newUser.matrixA)
	go func() {
		for i := 0; i < newUser.sizeMatrix; i++ {
			for j := 0; j < newUser.sizeMatrix; j++ {
				column := make([]float64, newUser.sizeMatrix)
				for k := 0; k < newUser.sizeMatrix; k++ {
					column[k] = newUser.matrixB[k][j]
				}
				jobChan <- job{i, j, &(newUser.matrixA[i]), &(column)}
			}
		}
	}()

	for i := 0; i < newUser.sizeMatrix*newUser.sizeMatrix; i++ {
		result := <-resultChan
		C[result.x][result.y] = result.value
	}

	resultMessage := matrixToString(C)
	elapsed := time.Since(begin)

	// Send result to client
	io.WriteString(newUser.connection, resultMessage+"$")
	newUser.connection.Close()
	fmt.Println("Client user disconnected", "time elapsed:", elapsed.String(), "Nb worker:", nb_worker)

	// Kill workers
	for k := 0; k < nb_worker; k++ {
		jobChan <- job{-1, -1, nil, nil}
	}
}

//  --------------------------------------------------
// Main function

func main() {
	port := os.Args[1]
	socket := host + ":" + port
	listen, err := net.Listen("tcp", socket)

	if err != nil {
		fmt.Println("Error while listening", err.Error())
		os.Exit(1)
	}

	fmt.Println("Server is listening on port", port)

	nbUsers := 0

	for {
		conn, err := listen.Accept()
		defer conn.Close()

		if err != nil {
			fmt.Println("Error while accepting connection", err.Error())
			os.Exit(1)
		}

		var newUser user
		newUser.connection = conn
		newUser.id = nbUsers

		go goHandlerUser(newUser)
		nbUsers++
	}

}
