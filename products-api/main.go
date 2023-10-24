package main

import (
	"fmt"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "    "
	dbname   = "productsdb"
)

func main() {
	a := App{}

	fmt.Printf(host)

	a.Initialize(host, user, password, dbname, port)

	a.Run(":8010")
}
