package main


// Struct for a database
// FIELDS:
//	name - Name of the database
// 	filePath - Absolute (i.e. from root) path to the JSON file (with '.json' suffix included) in which data is saved
//  indexes - Array of indexes in the DB
type Database struct { 
	FilePath string;
	// indexes []index;
}
