package connectDB

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func InitConnector() {
	// Replace the values with your PlanetScaleDB credentials
	dsn := os.Getenv("DSN")

	// Create a new database connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Successfully connected to PlanetScaleDB!")
}
