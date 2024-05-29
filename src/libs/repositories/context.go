package repositories

import (
	"database/sql"
	"log"
)

type DatabaseContext interface {
	Connect()
}

type AztebotDbContext struct {
	ConnectionString string
	SqlDb            *sql.DB
}

// type AztemarketDbContext struct {
// 	ConnectionString string
// 	SqlDb            *sql.DB
// }

func (c *AztebotDbContext) Connect() {

	db, err := sql.Open("mysql", c.ConnectionString)
	if err != nil {
		log.Fatal("Connection to database cannot be established :", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database at %s cannot be reached : %s", c.ConnectionString, err)
	}

	c.SqlDb = db

}
