package main

import (
	"fmt"
	"os"
	"strings"
)

// Error type for all command parser related errors
type parserError struct {
	message string
}

func (e *parserError) Error() string {
	return fmt.Sprintf("PARSER ERROR: %s", e.message)
}

func Parse(command string, coll *Collection) {
	tokens := strings.Split(command, " ")
	opcode := tokens[0]
	switch {

	case opcode == "createdb":
		// New DB name is first argument, rest are all new column names
		coll.NewDB(tokens[1], tokens[2:]...)

	case opcode == "dropdb":
		coll.DropDB(tokens[1])

	case opcode == "renamedb":
		coll.RenameDB(tokens[1], tokens[2])

	case opcode == "printcoll":
		coll.ListDBs()

	case opcode == "columns": // Print all columns of DB

		dbName := tokens[1]
		db, found_key := coll.DBs[dbName]
		// Raise non-fatal error & return from method if invalid database name provided
		if !found_key {
			err := parserError{fmt.Sprintf("Database %s does not exist in collection %s", dbName, coll.Name)}
			fmt.Println(err.Error())
			return
		}

		for _, name := range db.Columns {
			fmt.Println(name)
		}

	case opcode == "insert":
		dbName := tokens[1]
		db, found_key := coll.DBs[dbName]
		// Raise non-fatal error & return from method if invalid database name provided
		if !found_key {
			err := parserError{fmt.Sprintf("Database %s does not exist in collection %s", dbName, coll.Name)}
			fmt.Println(err.Error())
			return
		}

		// Parse columns & values from remaining command tokens
		columns := make([]string, 0, 10)
		values := make([]string, 0, 10)
		target := &columns
		for i := 2; i < len(tokens); i++ {
			if tokens[i] == "|" { // Pipe char seperates columns from values
				target = &values
				continue
			}

			*target = append(*target, tokens[i])
		}

		dbErr := db.Insert(columns, values)
		if dbErr != nil {
			fmt.Println(dbErr.Error())
		}

	case opcode == "exit":
		fmt.Println("Exiting...")
		os.Exit(0)

	default:
		fmt.Println("INVALID COMMAND: " + opcode)
	}
}
