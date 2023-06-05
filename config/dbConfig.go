package config

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GustavoFreitas22/ludus_sys/query"
	_ "github.com/lib/pq"
)

const postgresDriver = "postgres"

// TODO transform this in .env load data
const user = "postgres"

const password = "Postgres2023!"

const host = "localhost"

const port = "15432"

const dbName = "Ludus"

var Datasource *sql.DB

func CreateConnection() {
	log.Printf("Connecting %s ...", dbName)

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbName)

	datasource, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		log.Fatal("Error to connect: ", err)
	}

	log.Printf("Connected to %s!", dbName)

	Datasource = datasource

	initDatabase()
}

func initDatabase() {

	log.Println("Initialize database...")

	createTables, err := Datasource.Prepare(query.CREATE_TABLE_MODALITIES)
	checkErr(err)
	_, err = createTables.Exec()
	checkErr(err)

	createTables, err = Datasource.Prepare(query.CREATE_TABLE_REGISTRATION)
	checkErr(err)
	_, err = createTables.Exec()
	checkErr(err)

	createTables, err = Datasource.Prepare(query.CREATE_TABLE_PLAN)
	checkErr(err)
	_, err = createTables.Exec()
	checkErr(err)

	createTables, err = Datasource.Prepare(query.CREATE_TABLE_STUDY)
	checkErr(err)
	_, err = createTables.Exec()
	checkErr(err)

	createTables, err = Datasource.Prepare(query.CREATE_TABLE_PAYMENT)
	checkErr(err)
	_, err = createTables.Exec()
	checkErr(err)

	log.Println("Database initialize with success")

}

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Error to init database:", err)
	}
}
