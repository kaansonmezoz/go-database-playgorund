package service

import (
	manager "database/manager"
	"database/sql"
	"fmt"
	"log"
)

//TODO: The functions in the above should be parametrelized !

func CreateDatabase() bool {
	connection := manager.DatabaseManager.GetConnection()
	_, err := connection.Exec("CREATE DATABASE demoDB")

	if err == nil {
		log.Fatalf(err.Error())
		return false
	}

	fmt.Printf("Database has been created successfully !")

	return true
}

func CreateTable(createScript string) bool {
	connection := manager.DatabaseManager.GetConnection()
	statement, err := connection.Prepare(createScript)

	if err != nil {
		log.Fatalf(err.Error())
		return false
	}

	_, err = statement.Exec()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Table has been created successfully !")

	return true
}

func ExecuteSelectQuery(query string) *sql.Rows {
	connection := manager.DatabaseManager.GetConnection()
	rows, err := connection.Query(query)

	if err != nil {
		log.Fatalf(err.Error())
		return nil
	}

	fmt.Println("Select query has been executed successfully !")

	return rows
}
