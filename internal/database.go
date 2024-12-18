package internal

import (
	"fmt"
	"github.com/golang_db/internal/utils"
	"log"
	"math/rand/v2"
	"os"
)

// Database Struct for a database
// FIELDS:
//
//		FilePath - Absolute (i.e. from root) path to the CSV file (with '.csv' suffix included) in which data is saved
//	 Columns - In-order list of the names of the databases columns
//	 Indexes - Array of indexes in the DB
type Database struct {
	FilePath string
	Columns  []string
	// indexes []index;
}

// Error type for all db-related errors
type dbError struct {
	message string
}

func (e *dbError) Error() string {
	return fmt.Sprintf("DATABASE ERROR: %s", e.message)
}

// Insert Inserts a new entry into the DB, given some values and the columns they correspond to
// Returns a dbError if bad list of columns and values provided, or if we can't open or write to the database file
// Columns not specified in the parameters will be set to an empty cell for the new entry.
//
// PARAMS:
//
//	providedCols - a list of (user-provided) columns to add values for.
//	values - values[i] is the value to be added into the entry for column[i]
func (db *Database) Insert(providedCols []string, values []string) error {

	// Mismatch between columns and values
	if len(providedCols) != len(values) {
		return &dbError{fmt.Sprintf("Provided %d columns but %d values", len(providedCols), len(values))}
	}

	// Invalid column(s) listed
	columnsValid, invalidCol := utils.IsSubset(providedCols, db.Columns)
	if !columnsValid {
		return &dbError{fmt.Sprintf("Column '%s' does not exist in database", invalidCol)}
	}

	// Create CSV entry
	colValuesMap := utils.SlicesToMap(providedCols, values)
	entry := ""
	for _, col := range db.Columns {

		//  If col is a key in colValuesMap, the user has provided a value for this column
		value, valueProvided := colValuesMap[col]
		if valueProvided {
			entry += value
		}

		// Add an id for the entry
		if col == "id" {
			// Replace the random number w/ an auto-increment feature
			// with the counter stored in DB metadata file
			entry += fmt.Sprintf("%d", rand.IntN(100))
		}

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
		return &dbError{fmt.Sprintf("Couldn't open file %s", db.FilePath)}
	}
	defer file.Close()

	// Write entry to CSV file directly
	_, err = file.WriteString(entry)
	if err != nil {
		log.Fatal(err)
		return &dbError{fmt.Sprintf("Couldn't write to file %s", db.FilePath)}
	}

	return nil
}

// Select Returns some selected entries from a database that match a given condition string
//
// PARAMS: conditionStr - condition string
// RETURNS: A slice of strings that each represent an entry matching the condition
//
//	The strings in the slice are arranged in order of the columns to which they belong
//	(first string in slice will belong to first column, etc.)
/* func (db *Database) Select(conditionStr string) []string {

}
*/

// Update Updates column values of all entries from a database that match a given condition string
// PARAMS: conditionStr - condition string
func (db *Database) Update(conditionStr string) {

}

// Delete Deletes all entries from a database that match a given condition string
// PARAMS: conditionStr - condition string
func (db *Database) Delete(conditionStr string) {}
