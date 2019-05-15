package models

import (
	"fmt"

	"github.com/devmark88/unireg/config"
	"github.com/jinzhu/gorm"

	// import postgres dialectic
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init => Connect to Postgres Database
func Init(c *config.Specs) *gorm.DB {
	connStr := createConnectionString(c.DatabaseHost, c.DatabasePort, c.DatabasePassword, c.DatabaseUsername, c.DatabaseName)
	fmt.Println(connStr)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	err = db.DB().Ping()
	if err != nil {
		fmt.Println("cannot ping db")
		panic(err)
	}
	migrate(db)
	return db
}
func createConnectionString(host string, port int, pwd, username, dbname string) string {
	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s sslmode=disable", host, port, username, dbname, pwd)
}
func migrate(db *gorm.DB) {
	db.AutoMigrate(Graph{}, Node{}, Edge{})
	db.Model(Node{}).AddForeignKey("graph_id", "graphs(id)", "RESTRICT", "RESTRICT")
	db.Model(Node{}).AddForeignKey("node_id", "nodes(id)", "RESTRICT", "RESTRICT")
	db.Model(Edge{}).AddForeignKey("from", "nodes(id)", "RESTRICT", "RESTRICT")
	db.Model(Edge{}).AddForeignKey("to", "nodes(id)", "RESTRICT", "RESTRICT")
}
