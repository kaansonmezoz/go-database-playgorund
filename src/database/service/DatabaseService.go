package service

import(
	manager "database/manager"
	scripts "database/scripts"
	"fmt"
	"log"
) 


func CreateDatabase() bool{
	connection := manager.DatabaseManager.GetConnection()	
	_, err := connection.Exec("CREATE DATABASE demoDB")

	if err == nil {
		log.Fatalf(err.Error())
		return false
	}
	
	fmt.Printf("Database created successfully !")

	return true
}

func CreateTable() bool {
	connection := manager.DatabaseManager.GetConnection()
	statement, err := connection.Prepare(scripts.CreateEmployeeTable)

	if err != nil {
		log.Fatalf(err.Error())
		return false
	}

	_, err = statement.Exec()

	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println("Table created successfully !")
	
	return true
}