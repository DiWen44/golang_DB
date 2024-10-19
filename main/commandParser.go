package main

import  (
	"strings"
	"fmt"
	"os"
)


// Error type for all command parser related errors 
type parserError struct {
	message string;
}
 
func (e *parserError) Error() string{
	return fmt.Sprintf("PARSER ERROR: %s", e.message)
}


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

		case opcode == "insert"
			dbName := tokens[1]
			db, found_key := coll.DBs[dbName]
			// Raise non-fatal error & return from method if invalid database name provided
			if !found_key {
				err = parserError{ fmt.Sprintf("Database %s does not exist in collection %s", dbName, coll.Name) }
				fmt.Println(err.Error())
				return
			}

			dbErr := db.Insert()
			



		case opcode == "exit":
			fmt.Println("Exiting...")
			os.Exit(0)

		default:
			fmt.Println("INVALID COMMAND: " + opcode)
	}
}