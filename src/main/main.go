package main

import (
	"os"
	"bufio"
	"fmt"
	"log"
)


func main(){

	// Desired collection to open is provided as an OS arg
	collectionName := os.Args[1] 

	currentCollection, err := LoadCollection(collectionName)
	if err != nil { // If collection does not exist, make new one under that name
		currentCollection := MakeNewCollection(collectionName)
		fmt.Println("CREATED NEW COLLECTION: " + currentCollection.Name)
	} else {
		fmt.Println("LOADED COLLECTION: " + currentCollection.Name)
	}

	reader := bufio.NewReader(os.Stdin)

	for { // Command loop
		
		fmt.Printf("> ")
		command, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		Parse(command, currentCollection)
		fmt.Println() // Go to newline for next command
	
	}
}		