package database

import(
	"database/sql"
	"log"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()
	if err != nil{
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	host,port,user,password,dbname)

	db, err := sql.Open("postgres",psqlInfo)
	if err != nil{
		log.Fatal("Error connecting to database:",err)
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("Connot reach database:",err)
	}

	fmt.Println("Successfully connected to PostgreSQL!")

	DB = db
}