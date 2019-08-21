package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type databaseManager struct {
	OpenConnection  func()
	CloseConnection func()
	GetConnection   func() *sql.DB
}


var DatabaseManager = databaseManager{
	OpenConnection:  openConnection,
	CloseConnection: closeConnection,
	GetConnection:   getConnection,
}

// Designed to be long-lived.
// Don’t Open() and Close() databases frequently.
// Create one sql.DB object for each distinct datastore you need to access.
// Keep it until the program is done accessing that datastore.
// Pass it around as needed, or make it available somehow globally, but keep it open

var database *sql.DB

func openConnection() {
	if database != nil {
		return
	}

	var err error = nil
	var driverName = "mysql"
	var dataSource = "root:root@tcp(127.0.0.13306)/" // its value is driver spefic

	// It does not establish any connections to the database. It prepares the database abstraction for later use.
	// The first actual connection to the underlying datastore will be established lazily
	database, err = sql.Open(driverName, dataSource)

	if err != nil {
		log.Fatal(err)
	}

	// to check right away that the database is available and accessible
	err = database.Ping()

	if err != nil {
		log.Fatal(err)
	}
}

func closeConnection() {
	defer database.Close()
}

func getConnection() *sql.DB {
	return database
}
