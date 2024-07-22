package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

var testDB *sql.DB

var (
	dbUser     string
	dbPassword string
	dbDatabase string
	dbConn     string
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_TEST_NAME")
	dbConn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
}

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	cmd := exec.Command(
		"mysql",
		"-h",
		"127.0.0.1",
		"-u",
		dbUser,
		dbDatabase,
		"-e",
		"source ./testdata/setupDB.sql",
	)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command(
		"mysql",
		"-h",
		"127.0.0.1",
		"-u",
		dbUser,
		dbDatabase,
		"-e",
		"source ./testdata/cleanupDB.sql",
	)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}
