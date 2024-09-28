package main

import (
	"fmt"
	"os"
)


// Struct for a database
// FIELDS:
//	name - Name of the database
// 	filename - Name of JSON file (with '.json' suffix included) in which data is saved
//  indexes - Array of indexes in the DB
type Database struct { 
	name string;
	filename string;
	// indexes []index;
}


// Creates a new database, file 
// and everything
func MakeNewDB(name string) *Database {
	filename := name + ".json"
	file, err := os.Create()
	file.close()

	res := &Database{name, filename}
	return res
}


// Loads a database from an existing file
func LoadDB(name string) *Database {

	filename := name + ".json"


	res := &Database{name, filename}
	return res
}


// Deletes a database
func DropDB(name string) {
	
}