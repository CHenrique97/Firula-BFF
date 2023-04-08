package connectDB

import (
	"database/sql"
	"log"
	"os"

	"github.com/firula-bff/initializers"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	initializers.LoadEnv()
}
func InitConnector() {
	// Replace the values with your PlanetScaleDB credentials
	log.Println(os.Getenv("DSN"))
	db, err := sql.Open("mysql", os.Getenv("DSN"))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping: %v", err)
	}
	log.Println("Successfully connected to PlanetScale!")
}
