package main

import (
	"log"
	"os"
)

func tryOpenFile() {
	f, err := os.OpenFile("notes.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}

/*
func main() {
	fmt.Println("this is simple function, hello world.")
	//tryOpenFile()

	stringMap := map[int]string{1: "A", 2: "B", 3: "C"}
	for k, v := range stringMap {
		fmt.Printf("%d: %s\n", k, v)
	}
}
*/
