package main

import (
	"architecture_go/pkg/store/postgres"
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	db := postgres.DBconnection("postgres", "123456789", "localhost", "5432", "Company")

	fmt.Println("DB connection: ", db)

	defer db.Close()
}
