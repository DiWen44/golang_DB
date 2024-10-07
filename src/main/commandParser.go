package main

import  (
	"strings"
	"fmt"
	"os"
)


func Parse(command string, coll *Collection) {
	tokens := strings.Split(command, " ")
	opcode := tokens[0]
	switch opcode {

		case "createdb":
			coll.NewDB(tokens[1])

		case "dropdb":
			coll.DropDB(tokens[1])

		case "renamedb":
			coll.RenameDB(tokens[1], tokens[2])

		case "printcoll":
			coll.ListDBs()

		case "exit":
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("INVALID COMMAND: " + opcode)
	}
}