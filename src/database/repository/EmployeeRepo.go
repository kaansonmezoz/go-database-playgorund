package repository

import (
	models "database/models"
	scripts "database/scripts"
	service "database/service"
	"errors"
	"log"
)

func CreateTable() {
	isSuccessful := service.CreateTable(scripts.CreateEmployeeTable)

	if isSuccessful {
		// do sth
	} else {
		// do sth
	}
}

func GetEmployeeById(id int) (models.EmployeeModel, error) {
	var employee *models.EmployeeModel

	rows := service.ExecuteSelectQuery(scripts.SelectEmployeeByIdScript, id)

	if rows == nil {
		return *employee, errors.New("EMPTY_QUERY_RESULT") // TODO: Bunları bir error tiplerimiz olsun onlar icersine gomerek yapabiliriz aslinda
	}

	defer rows.Close()

	employee = new(models.EmployeeModel)

	for rows.Next() {
		err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName)

		if err != nil {
			log.Fatal(err)
			return *employee, errors.New("ROW_SCAN_ERROR")
		}

		log.Println("Query Result: %#v", employee)
	}

	return *employee, nil
}

func GetAllEmployees() ([]models.EmployeeModel, error) {
	var employees []models.EmployeeModel

	rows := service.ExecuteSelectQuery(scripts.SelectAllEmployeesScript, nil)

	if rows == nil {
		return employees, errors.New("EMPTY_QUERY_RESULT")
	}

	for rows.Next() { // TODO: burası da bir methoda extract edilebilir
		var employee = new(models.EmployeeModel)

		err := rows.Scan(&employee.Id, &employee.FirstName, &employee.LastName)

		if err != nil {
			log.Fatal(err)
			return employees, errors.New("ROW_SCAN_ERROR")
		}

		log.Println("Query Result: %#v", employee)
		employees = append(employees, *employee)
	}

	return employees, nil
}
