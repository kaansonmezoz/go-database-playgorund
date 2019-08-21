package repository

import (
	scripts "database/scripts"
	service "database/service"
)

func CreateTable() {
	isSuccessful := service.CreateTable(scripts.CreateEmployeeTable)

	if isSuccessful {
		// do sth
	} else {
		// do sth
	}
}
