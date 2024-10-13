package main

import  (
	"strings"
	"fmt"
	"os"
)


func Parse(command string, coll *Collection) {
	tokens := strings.Split(command, " ")
	opcode := tokens[0]
	switch {

		case opcode == "createdb":
			coll.NewDB(tokens[1])

		case opcode == "dropdb":
			coll.DropDB(tokens[1])

		case opcode == "renamedb":
			coll.RenameDB(tokens[1], tokens[2])

		case opcode == "printcoll":
			coll.ListDBs()

		case opcode == "exit":
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("INVALID COMMAND: " + opcode)
	}
}