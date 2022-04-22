package main

import (
	"CourierManagement/DataContext"
)

//***Note**
// Blank import identifier: https://go.dev/ref/spec#Import_declarations
//-----------------
//Install drivers
// go get -u gorm.io/gorm
// go get -u gorm.io/driver/sqlserver
//------------
//Find SQL server port
// 	USE master
// GO
// xp_readerrorlog 0, 1, N'Server is listening on'
// GO
func main() {

	DataContext.CallDB()
	// dbContext.CallDB()
}
