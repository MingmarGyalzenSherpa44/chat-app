package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

var DBConn *pgx.Conn

func Connect() (*pgx.Conn, error) {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DBUsername := os.Getenv("DB_USERNAME")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")

	conn, err := pgx.Connect(context.Background(), fmt.Sprintf("postgres://%v:%v@localhost:5432/%v", DBUsername, DBPassword, DBName))
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func init() {
	var err error

	DBConn, err = Connect()
	if err != nil {
		log.Fatal(err)
	}

}

func RegisterUser(email, username, password string) error {

	_, err := DBConn.Exec(context.Background(), "INSERT INTO users (email, username, password) VALUES ($1, $2, $3)", email, username, password)
	return err

}

func LoginUser(username, password string) error {

	type userCredential struct {
		Username string
		Email    string
		Password string
	}
	dbCredential := userCredential{}

	//fetch data
	err := DBConn.QueryRow(context.Background(), "SELECT email, username, password FROM users WHERE username = $1", username).Scan(&dbCredential.Email, &dbCredential.Username, &dbCredential.Password)

	if err := bcrypt.CompareHashAndPassword([]byte(dbCredential.Password), []byte(password)); err != nil {
		log.Fatal("Username or Password is incorrect!")
	}

	fmt.Printf("Login successfully!!, Welcome %v \n", username)

	return err
}
