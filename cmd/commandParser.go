package cmd

import (
	"fmt"
	"github.com/golang_db/internal"
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

// Used to check if adequate number of tokens/arguments provided for command
//
// PARAMS:
//
//	expected - how many args were expected
//	args - array of arguments (not including command opcode itself)
//
// RETURNS:
//
//	parser error w/ appropriate message if wrong no. of args
//	otherwise nil
func errorIfUnexpectedNumArgs(expected int, args []string) *parserError {
	if len(args) != expected {
		return &parserError{fmt.Sprintf("EXPECTED %d ARGUMENTS, GOT %d", expected, len(args))}
	}
	return nil
}

func Parse(command string, coll *internal.Collection) {
	tokens := strings.Split(command, " ")
	opcode := tokens[0]
	args := tokens[1:]
	switch {

	case opcode == "createdb":
		// New DB name is first argument, rest are all new column names
		coll.NewDB(args[0], args[1:]...)

	case opcode == "dropdb":
		err := errorIfUnexpectedNumArgs(1, args)
		if err != nil {
			fmt.Println(err.Error())
		}

		err2 := coll.DropDB(args[0])
		if err2 != nil {
			fmt.Println(err2.Error()) // Display any errors passed forward by dropDB
		}

	case opcode == "renamedb":
		err := errorIfUnexpectedNumArgs(2, args)
		if err != nil {
			fmt.Println(err.Error())
		}

		err2 := coll.RenameDB(args[0], args[1])
		if err2 != nil {
			fmt.Println(err2.Error())
		}

	case opcode == "listdbs":
		coll.ListDBs()

	case opcode == "columns": // Print all columns of DB

		err := errorIfUnexpectedNumArgs(1, args)
		if err != nil {
			fmt.Println(err.Error())
		}

		dbName := args[0]
		db, foundKey := coll.DBs[dbName]
		// Raise non-fatal error & return from method if invalid database name provided
		if !foundKey {
			err2 := internal.collError{fmt.Sprintf("NO DATABASE CALLED '%s' IN COLLECTION '%s'", dbName, coll.Name)}
			fmt.Println(err2.Error())
			return
		}

		for _, name := range db.Columns {
			fmt.Println(name)
		}

	case opcode == "insert":
		dbName := args[0]
		db, foundKey := coll.DBs[dbName]
		// Raise non-fatal error & return from method if invalid database name provided
		if !foundKey {
			err := internal.collError{fmt.Sprintf("NO DATABASE CALLED '%s' IN COLLECTION '%s'", dbName, coll.Name)}
			fmt.Println(err.Error())
			return
		}

		// Parse columns & values from remaining arguments
		columns := make([]string, 0, 10)
		values := make([]string, 0, 10)
		target := &columns
		for i := 1; i < len(args); i++ {
			if args[i] == "|" { // Pipe char seperates columns from values
				target = &values
				continue
			}
			*target = append(*target, args[i])
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
