package main

import (
	"bufio"
	"fmt"
	"github.com/golang_db/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

// Collection Represents a collection (a group of databases)
// Used to hold the currently active collection that the user is operating on
// In the filesystem, a collection is a directory that holds JSON files, each of them representing a database
//
// FIELDS:
//
//		name - name of the collection
//		path - file path of the directory associated w/ the collection
//	 dbs - map of databases in collection: key is database name, value is a pointer to database object
type Collection struct {
	Name string
	Path string
	DBs  map[string]*database.Database
}

// Error type for all collection-related errors
type collError struct {
	message string
}

func (e *collError) Error() string {
	return fmt.Sprintf("COLLECTION ERROR: %s", e.message)
}

// LoadCollection Loads an existant collection from the filesystem
// If collection of specified name does not exist, returns a collectionError
//
// PARAMS:
//
//	name - name of the collection
func LoadCollection(name string) (*Collection, error) {

	// Get collection directory
	// Concatenate collection name onto .env variable for collections directory path
	godotenv.Load("/Users/devinsidhu/Documents/golang_db/.env")
	collectionPath := fmt.Sprintf("%s/%s", os.Getenv("COLLECTIONS_DIR"), name)
	collectionDir, err := os.Open(collectionPath)
	if err != nil {
		return nil, err
	}

	// Get database files from collection directory
	filenames, err := collectionDir.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}
	err = collectionDir.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Load databases into dbs map
	dbs := make(map[string]*database.Database)
	for _, filename := range filenames {
		dbName := filename[:len(filename)-4] // Remove last 5 characters (i.e. '.json' extension) from filename to get database's name
		dbFilePath := fmt.Sprintf("%s/%s", collectionPath, filename)
		dbs[dbName] = loadDB(dbFilePath)
	}

	return &Collection{Name: name, Path: collectionPath, DBs: dbs}, nil
}

// MakeNewCollection Makes an entirely new collection
// PARAMS: name - name of the new collection
func MakeNewCollection(name string) *Collection {

	// Make collection directory
	// Concatenate collection name onto .env variable for collections directory path
	godotenv.Load("/Users/devinsidhu/Documents/golang_db/.env")
	collection_path := fmt.Sprintf("%s/%s", os.Getenv("COLLECTIONS_DIR"), name)
	fmt.Println(collection_path)
	err := os.Mkdir(collection_path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Empty slice of dbs, since collection is new
	dbs := make(map[string]*database.Database)
	return &Collection{Name: name, Path: collection_path, DBs: dbs}
}

// NewDB Creates a new database in the filesystem and add it to the collection
// PARAMS:
//
//	DBName - name of new DB
//	columns - names of new columns for DB (variadic, so can provide 1 slice of strings, or all strings as separate arguments)
func (coll *Collection) NewDB(DBName string, columns ...string) {

	// Add an ID column as first column in DB
	columns = append([]string{"id"}, columns...)

	// Create JSON file for DB
	DBPath := fmt.Sprintf("%s/%s.csv", coll.Path, DBName)
	file, err := os.Create(DBPath)
	if err != nil {
		log.Fatal(err)
	}

	// Add columns to first line of newly created CSV file
	columnsStr := strings.Join(columns, ",") + "\n"
	_, err = file.WriteString(columnsStr)
	if err != nil {
		log.Fatal(err)
	}

	closeErr := file.Close()
	if closeErr != nil {
		log.Fatal(closeErr)
	}

	// Add DB to active collection
	coll.DBs[DBName] = &database.Database{FilePath: DBPath, Columns: columns}
}

// Loads a database from an existing file into a database object
// Returns a pointer to the new database object
// This is to be called when loading an existing collection on program startup
//
// Unlike NewDB(), this is a standalone function (not a collection struct method)
// and so doesn't add the DB to the collection's DB map
//
// PARAMS: filePath - path to JSON file associated with DB to load
func loadDB(filePath string) *database.Database {
	// Load DB file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lineScanner := bufio.NewScanner(file)

	// Read CSV columns from 1st line of file
	lineScanner.Scan()
	columnStr := lineScanner.Text()
	columns := strings.Split(columnStr, ",")

	res := &database.Database{FilePath: filePath, Columns: columns}
	return res
}

// DropDB Drops a database, removing it from the collection and deleting it's JSON file from the filesystem
// Returns a collection Error
//
// PARAMS:
//
//	dbName - name of DB to drop
func (coll *Collection) DropDB(dbName string) error {

	// Delete DB file
	filepath := fmt.Sprintf("%s/%s.csv", coll.Path, dbName)
	err := os.Remove(filepath)
	if err != nil {
		return &collError{fmt.Sprintf("NO DATABASE CALLED '%s' IN COLLECTION '%s'", dbName, coll.Name)}
	}

	// Remove DB from collection object's DB map
	delete(coll.DBs, dbName)
	return nil
}

// RenameDB Rename a database in the collection
//
// PARAMS:
//
//	oldDBName - current name of DB to rename
//	newDBName - New name for DB
func (coll *Collection) RenameDB(oldDBName string, newDBName string) error {

	// Rename DB in collection by adding new pair under new name and deleting old entry
	db, foundKey := coll.DBs[oldDBName]
	if !foundKey {
		return &collError{fmt.Sprintf("NO DATABASE CALLED '%s' IN COLLECTION '%s'", oldDBName, coll.Name)}
	}
	coll.DBs[newDBName] = db
	delete(coll.DBs, oldDBName)

	// Rename DB file
	oldPath := db.FilePath
	newPath := fmt.Sprintf("%s/%s.csv", coll.Path, newDBName)
	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

// ListDBs Outputs a list of all databases in the collection
func (coll *Collection) ListDBs() {
	for name, _ := range coll.DBs {
		fmt.Println(name)
	}
}
