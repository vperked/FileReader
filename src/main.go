package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	openFile()
}

func openFile() {
	fileLocation := "/FileReader/data/retard.txt"
	file, err := os.Open(fileLocation)
	if err != nil {
		fmt.Println("Couldnt open file!", err)
		defer file.Close()

		dat, err := os.ReadFile(fileLocation)
		if err != nil {
			fmt.Println("Couldnt read file!", err)
			fmt.Print(string(dat))

			data, err := io.ReadAll(file)
			if err != nil {
				fmt.Println("Couldnt read the byte data.", err)
			}
			fmt.Println(string(data))
		}
	}
}
