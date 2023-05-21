// CLIENT PART //

package main

//  --------------------------------------------------
// imports
import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

//  --------------------------------------------------
// Global variables

var host string = "localhost"
var port string

//  --------------------------------------------------
// Functions

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func write(text string, file *os.File) {
	_, err := file.WriteString(text)
	check(err)
}

func read(filename string) string {
	data, err := ioutil.ReadFile(filename)
	check(err)
	return string(data)
}

//  --------------------------------------------------
// Main function

func main() {
	if len(os.Args) != 3 {
		println("Usage: go run client.go <port> <inputPath>")
		os.Exit(1)
	}

	cwd, _ := os.Getwd()
	port = os.Args[1]
	inputPath := os.Args[2]
	socket := host + ":" + port

	if _, err := os.Stat(inputPath); errors.Is(err, os.ErrNotExist) {
		println("Input file not found in " + cwd + "/" + inputPath)
		os.Exit(1)
	}
	inputFile, err := os.OpenFile(inputPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	defer inputFile.Close()
	check(err)

	connection, err := net.Dial("tcp", socket)
	defer connection.Close()
	check(err)

	data := read(inputFile.Name()) + "$"
	io.WriteString(connection, data)
	inputFile.Close()

	reader := bufio.NewReader(connection)
	message, err := reader.ReadString('$')
	check(err)
	message = strings.TrimSuffix(message, "$")

	outputFile, err := os.OpenFile("output.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	defer outputFile.Close()
	check(err)
	write(message, outputFile)
	outputFile.Close()

	connection.Close()
	fmt.Println("File saved to " + "output.txt")
}
