package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//Config Database
const (
	username = "root"
	password = "1234"
	hostname = "localhost:3306"
	dbname   = "test"
)

func dsn(dbName string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func main() {

	db, err := sql.Open("mysql", dsn(dbname))
	if err != nil {
		log.Printf("Error %s when opening DB\n", err)
	}
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
		return
	}
	log.Printf("Connected to DB successfully\n")

	_, err = db.ExecContext(ctx, `CREATE TABLE Product (
		id INT NOT NULL AUTO_INCREMENT,
		product_code VARCHAR(45) NOT NULL,
		product_name VARCHAR(45) NOT NULL,
		quantity INT NOT NULL,
		PRIMARY KEY (id));
	  `)
	if err != nil {
		log.Printf("Error %s when creating Table\n", err)
		return
	}
	log.Printf("Create table successfully\n")

	return

}
