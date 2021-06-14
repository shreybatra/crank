package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func printOutput(response interface{}) {
	colorReset := "\033[0m"
	colorGreen := "\033[32m"
	fmt.Println(string(colorGreen), response, colorReset)
}

func main() {

	conn, err := net.Dial("tcp", "localhost:9876")
	if err != nil {
		fmt.Println(err)
		return
	}

	var file *os.File

	if len(os.Args) > 1 {
		fileName := os.Args[1]

		re := regexp.MustCompile(`^[a-zA-Z0-9]+\.gsb`)

		if !re.MatchString(fileName) {
			log.Fatalf("Wrong file / filename.")
		}

		file, err = os.Open(fileName)
		if err != nil {
			log.Fatalf("Cannot read file")
		}

	} else {
		file = os.Stdin
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	defer conn.Close()

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

			conn.Write([]byte(strings.TrimRight(message, "\n")))

			message := make([]byte, 4096)
			length, err := conn.Read(message)

			if err != nil {
				fmt.Println("Connection closed by server.")
				break
			}

			if length > 0 {
				printOutput(string(message))
			}
		}
		fmt.Print("> ")
	}

}
