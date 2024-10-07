package main

import (
	"fmt"
)


// Struct for a database
// FIELDS:
//	name - Name of the database
// 	filePath - Absolute (i.e. from root) path to the JSON file (with '.json' suffix included) in which data is saved
//  indexes - Array of indexes in the DB
type Database struct { 
	filePath string;
	// indexes []index;
}




// Error type for database errors
// 
// FIELDS:
//  name - name of the database
//  message - Error message
type databaseError struct {
	name string
	message string
}


// databaseError's implementation of Error()
// This exists to allow databaseError to satisfy the builtin error interface
func (e *databaseError) Error() string {
	return fmt.Sprintf("'%s' -- %s", e.name, e.message)
}