package main

import (
	"os"
	"fmt"
	"log"
)

// Struct for a database
// FIELDS:
// 	FilePath - Absolute (i.e. from root) path to the CSV file (with '.csv' suffix included) in which data is saved
//  Columns - In-order list of the names of the databases columns
//  Indexes - Array of indexes in the DB
type Database struct { 
	FilePath string;
	Columns []string;
	// indexes []index;
}


// Error type for all db-related errors 
type dbError struct {
	message string;
}
 
func (e *dbError) Error() string{
	return fmt.Sprintf("DATABASE ERROR: %s", e.message)
}


// Inserts a new entry into the DB, given some values and the columns they correspond to
// Returns a dbError if bad list of columns and values provided, or if can't open or write to the database file
// Columns not specified in the parameters will be set to an empty cell for the new entry.
//
// PARAMS:
//	providedCols - a list of (user-provided) columns to add values for.
//	values - values[i] is the value to be added into the entry for column[i]
func (db *Database) Insert(providedCols []string, values []string) error {

	// Mismatch between columns and values
	if len(providedCols) != len(values){
		return &dbError{ fmt.Sprintf("Provided %d columns but %d values", len(providedCols), len(values)) }
	}

	// Invalid column(s) listed
	columnsValid, invalidCol := isSubset(providedCols, db.Columns)
	if !columnsValid {
		return &dbError{ fmt.Sprintf("Column '%s' does not exist in database", invalidCol) }
	}

	// Create CSV entry
	colValuesMap := slicesToMap(providedCols, values)
	entry := ""
	for _, col := range db.Columns {

		// User has provided a value for this column if col in colValuesMap
		value, value_provided := colValuesMap[col]
		if value_provided { entry += value }
		
		// If cell not at last column, add comma to end, Otherwise add newline
		if col != db.Columns[len(db.Columns)-1] {  
			entry += ","
		} else { 
			entry += "\n"
		}
	}

	// Open db file
	file, err := os.OpenFile(db.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return &dbError{ fmt.Sprintf("Couldn't open file %s", db.FilePath) }
	}
	defer file.Close()

	// Write entry to CSV file directly
	_, err = file.WriteString(entry)
	if err != nil {
		log.Fatal(err)
		return &dbError{ fmt.Sprintf("Couldn't write to file %s", db.FilePath) }
	}
		
	return nil
}


// Private utility function to check if a slice of strings is a subset of another.
// If is a subset, returns true and an empty string
// If not a subset, returns false along with the first string found not to be in the superset
//
// PARAMS:
// 	a1 - smaller slice (the possible subset)
//	a2 - bigger slice (the possible superset)
//
// RETURNS:
//  bool - is a1 a subset?
//	string - first string in a1 that is found to not be in a2 (if a1 found to be a subset, this is nil)
func isSubset(a1 []string, a2 []string) (bool, string) {

	// Make set-type structure from a2
	a2Set := make(map[string]bool)
	for _, s2 := range a2{
		a2Set[s2] = true
	}

	// Check if all strings in a1 are present in a2
	for _, s1 := range a1{
		// If a key is not present in the map
		// Golang initializes that key's value to the zero value of it's type
		// e.g. for bool-type values, this will be false
		if a2Set[s1] != true {
			return false, s1
		}
	}

	return true, ""
}


// Private utility function
// Takes 2 slices (of equal length) that contain strings
// returns a map derived from taking the first slice's elements as keys and the second slice's as values.
// Such that map[a1[i]] = a2[i] where 0 < i < n.
//
// PARAMS:
//	a1 - Slice from which to get keys of map
//	a2 - Slice from which to get values of map
func slicesToMap(a1 []string, a2 []string) map[string]string {

	res := make(map[string]string)
	for i := 0; i < len(a1); i++{
		res[a1[i]] = a2[i]
	}

	return res
}


