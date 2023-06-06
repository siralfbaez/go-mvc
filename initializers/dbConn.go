package initializers

import (
	_ "database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/siralfbaez/go-mvc/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB

func ConnToDb() {
	var err error
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	var (
		ssuser string
		dbname string
	)
	DB.Raw("SELECT session_user;").Scan(&ssuser)
	DB.Raw("SELECT current_database();").Scan(&dbname)

	if err != nil {
		fmt.Printf("\n ====================================================================\n")
		defer log.Fatalf("   User:=> " + ssuser + " :-( Failed to connect to" + dbname + "database!\n")
		fmt.Printf(" ====================================================================\n")
	} else {
		fmt.Printf("\n =======================================================================\n")
		fmt.Printf("  User:=> " + ssuser + " :-) successfully connected to:=> " + dbname + " <=:database!\n")
		fmt.Printf(" =======================================================================\n")
	}
}

func SyncDB() {
	err := DB.AutoMigrate(&models.Post{})
	if err != nil {
		return
	}
}
