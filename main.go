package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// TODO custom flags for variables
	host := os.Getenv("DOCKER_MACHINE_IP")
	username := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	// dbport := os.Getenv("POSTGRES_PORT")
	// if dbport == "" {
	// 	dbport = "5432"
	// }
	dbname := os.Getenv("POSTGRES_NAME")

	log.Println("Running dbrefresh")
	postgresInfo := fmt.Sprintf("%s://%s:%s@%s/postgres?sslmode=disable", dbname, username, password, host)
	// postgresInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, dbport)
	db, err := sql.Open("postgres", postgresInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Dropping public SCHEMA")
	_, err = db.Query("DROP SCHEMA public CASCADE;")
	if err != nil {
		log.Fatal("failed to drop public schema", err)
	}

	log.Println("Creating public SCHEMA")
	_, err = db.Query("CREATE SCHEMA public;")
	if err != nil {
		log.Fatal("failed to create public schema")
	}

	log.Println("done.")
}
