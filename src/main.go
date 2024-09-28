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
		fmt.Println("CREATED NEW COLLECTION: " + collectionName)
	} else {
		fmt.Println("LOADED COLLECTION: " + collectionName)
	}

	var command string
	reader := bufio.NewReader(os.Stdin)

	for { // Command loop
		
		fmt.Printf("> ")
		command, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if command == "quit" {
			fmt.Println("Exiting...")
			os.Exit(0)
		}

		print(command)
		fmt.Println() // Go to newline for next command
	
	}

}		