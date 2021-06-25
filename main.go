package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/shreybatra/gocrank"
)

func printOutput(response interface{}, err error) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	colorRed := "\u001b[31m"

	if err != nil {
		fmt.Println(string(colorRed), err, colorReset)
		return
	}

	fmt.Println(string(colorGreen), response, colorReset)
}

func extractNext(message string) (string, string) {
	msgLen := len(message)
	var i int
	for i = 0; i < msgLen; i++ {
		if message[i] == ' ' {
			break
		}
	}

	command := message[0:i]

	if i < msgLen {
		i++
	}

	return command, message[i:]
}

func handleSetCommand(conn *gocrank.GoCrank, message string) {
	key, message := extractNext(message)
	var data interface{}
	json.Unmarshal([]byte(message), &data)
	result, err := conn.Set(key, data)
	printOutput(result, err)
}

func handleGetCommand(conn *gocrank.GoCrank, message string) {
	key, _ := extractNext(message)
	result, err := conn.Get(key)
	printOutput(result, err)
}

func handleFindCommand(conn *gocrank.GoCrank, message string) {
	var data interface{}
	json.Unmarshal([]byte(message), &data)
	switch data := data.(type) {
	case map[string]interface{}:
		result, err := conn.Find(data)
		printOutput(result, err)
	default:
		result, err := "", errors.New("invalid arguments to find command")
		printOutput(result, err)
	}
}

func main() {

	conn := gocrank.NewCrankConnection("localhost:9876")

	var file *os.File

	if len(os.Args) > 1 {
		fileName := os.Args[1]

		// Support for reading .gsb text files
		re := regexp.MustCompile(`^[a-zA-Z0-9]+\.gsb`)

		if !re.MatchString(fileName) {
			log.Fatalf("Wrong file / filename.")
		}

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatalf("Cannot read file")
		}
		file = f

	} else {
		file = os.Stdin
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	fmt.Print("> ")
	for scanner.Scan() {
		message := scanner.Text()
		if file != os.Stdin {
			fmt.Println(message)
		}

		if strings.TrimRight(message, "\n") != "" {

			if message == "exit" {
				os.Exit(0)
			}

			command, message := extractNext(message)

			switch command {
			case "set":
				handleSetCommand(conn, message)
			case "get":
				handleGetCommand(conn, message)
			case "find":
				handleFindCommand(conn, message)
			default:
				result, err := "", errors.New("wrong command")
				printOutput(result, err)
			}

		}
		fmt.Print("> ")
	}

}
