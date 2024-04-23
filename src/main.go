package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Enter a valid file path : ")
	openFile()

}

func openFile() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	fmt.Println("Reading:", scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	dat, err := os.ReadFile(scanner.Text())
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(dat))
	}
	file, err := os.Open(scanner.Text())
	if err != nil {
		fmt.Println(err)
	}
	b1 := make([]byte, 60)
	n1, err := file.Read(b1)
	if err != nil {
		fmt.Println("Couldnt read bytes.", err)
	}
	fmt.Println(string(b1[:n1]))
}
