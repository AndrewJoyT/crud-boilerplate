package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/AndrewJoyT/crud-boilerplate/database/schema"
)

var (
	Client   *sql.DB
	Database *gorm.DB
)

func Connect() {
	host := os.Getenv("CONFIGDB_HOST")
	port := os.Getenv("CONFIGDB_PORT")
	user := os.Getenv("CONFIGDB_USER")
	password := os.Getenv("CONFIGDB_PASS")
	dbname := os.Getenv("CONFIGDB_DBNAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
	fmt.Println(conn)
	db, err := gorm.Open("mysql", conn)
	if err != nil {
		log.Fatal("database not connected :", err)
	}
	db.LogMode(true)

	Database = db
	Client = db.DB()

	autoCreate := os.Getenv("CONFIGDB_AUTO_CREATE")
	if autoCreate == "true" {
		fmt.Println("Dropping and recreating all tables...")
		schema.AutoMigrate(db)
		fmt.Println("All tables recreated successfully...")
	}

	fmt.Printf("\ndatabase succesfully connected. auto create: %s\n", autoCreate)
}
