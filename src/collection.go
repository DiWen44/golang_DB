package main

import  (
	"errors"
	"os"
	"fmt"
	"log"
)


// Represents a collection (a group of databases)
// Used to hold the currently active collection that the user is operating on
// In the filesystem, a collection is a directory that holds JSON files, each of them representing a database
// 
// FIELDS:
//	name - name of the collection
//  dbs - slice of pointers to databases in the collection 
type Collection struct {
	name string
	dbs []*database
}


// Loads an existant collection from the filesystem
//
// PARAMS:
// 	name - name of the collection
func LoadCollection(name string) (*Collection, error) {
	// Get collection directory
	collection_dir, err := os.Open("/"+name)
	if err != nil {
		return nil, &collectionError{name, fmt.Sprintf("Could not open collection '%s'", name)}
	}

	// Get database files from collection directory
	files, err := collection_dir.Readdir(0)
	if err != nil {
        log.Fatal(err)
    }

    collection_dir.Close()
    if err != nil {
        log.Fatal(err)
    }

    // Load databases in 
    var dbs []*database 
    for file := range files{
    	dbs = append( dbs, LoadDB(file.Name()) )
    }

    return &Collection{name, dbs}, nil
}


// Makes an entirely new collection
//
// PARAMS:
// 	name - name of the collection
func MakeNewCollection(name string) *collection {

	// Make collection directory
	err := os.Mkdir(name, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Empty slice of dbs, since collection is new
	var dbs []*database = {}
	return &Collection{name, dbs}
}


// Error type for collection errors
// 
// FIELDS:
//  name - name of the collection
//  message - Error message
type collectionError struct {
	name string
	message string
}

// collectionError's implementation of Error()
// This exists to allow collectionError to satisfy the builtin error interface
func (e *commandError) Error() string {
	return fmt.Sprintf("'%s' -- %s", e.section, e.message)
}


