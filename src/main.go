package main

import (
	"fmt"
	"os"
)

func main() {
	openFile()
}

func openFile() {
	fileLocation := "/Users/Perke/OneDrive/Desktop/FileReader/data/retard.json"
	dat, err := os.ReadFile(fileLocation)
	if err != nil {
		fmt.Println(err)
		fmt.Println(string(dat))
	}
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println(err)
	}
	b1 := make([]byte, 13)
	n1, err := file.Read(b1)
	if err != nil {
		fmt.Println("Couldnt read bytes.", err)
	}
	fmt.Printf("%d bytes : %s\n", n1, string(b1[:n1]))
}
