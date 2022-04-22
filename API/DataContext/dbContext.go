package DataContext

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	_ "gorm.io/driver/sqlserver"
	_ "gorm.io/gorm"
)

// "context"
// "database/sql"
// "log"

// _ "github.com/denisenkom/go-mssqldb"
// _ "gorm.io/driver/sqlserver"
// _ "gorm.io/gorm"

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
func CallDB() {
	// return "Hello test"
	var ctx = context.Background()
	var conn = "sqlserver://GoTest:gotest@localhost:49677?database=Delivery&connection+timeout=30"
	db, err := sql.Open("sqlserver", conn)
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal("Erro in conection: ", err)
	}
	log.Printf("connected !\n")
	//------Query DB

	rows, err := db.Query("SELECT ProductName FROM [Order]where Id=1")
	if err != nil {
		log.Fatal("Erro in Query: ", err)
	}
	if rows == nil {
		log.Fatal("No record found ")
		return
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			// Check for a scan error.
			// Query rows will be closed with defer.
			log.Fatal(err)
		}
		log.Printf("P: %s", name)
	}
	if err != nil {
		log.Fatal("Error creating connection pool: " + err.Error())
	}

	defer db.Close()
}
