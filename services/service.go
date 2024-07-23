package services

import "database/sql"

type MyAppService struct {
	db *sql.DB
}

func NewMyAppSrvice(db *sql.DB) *MyAppService {
	return &MyAppService{db: db}
}
