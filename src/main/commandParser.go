package main

import  (
	"strings"
	"fmt"
)


// Error type for command errors picked up in parse()
// 
// FIELDS:
//  section - the section (part of the command) in question that caused the error
//  message - Error message
type commandError struct {
	section string
	message string
}


// commandError's implementation of Error()
// This exists to allow commandError to satisfy the builtin error interface
func (e *commandError) Error() string {
	return fmt.Sprintf("'%s' -- %s", e.section, e.message)
}


func parse(command string) error {
	tokens := strings.Split(command, " ")
	switch opcode := tokens[0]; opcode {

		case "createdb":

		case "dropdb":

		case "renamedb":

		case "insert":

		case "update":

		case "select":

		case "delete":

		case "createind":

		case "deleteind":

		default:
			return &commandError{opcode, "INVALID OPERATION"}

	}

	return nil // If no errors were encountered

}