package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mohammadinasab-dev/employee-salary/rest"
)

func main() {
	log.Fatal(rest.RunApi("."))
}
