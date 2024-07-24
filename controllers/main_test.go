package controllers_test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/S-Yamamoto3215/go-myapi/controllers"
	"github.com/S-Yamamoto3215/go-myapi/services"
	"github.com/joho/godotenv"
)

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

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	testDB, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Println("DB setup fail")
		os.Exit(1)
	}

	ser := services.NewMyAppService(testDB)
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
