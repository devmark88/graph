package models

import (
	"fmt"

	"github.com/devmark88/unireg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Init => Connect to Postgres Database
func Init(c *config.Specs) *gorm.DB {
	connStr := createConnectionString(c.DatabaseHost, c.Port, c.DatabasePassword, c.DatabaseUsername, c.DatabaseName)
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}
	err = db.DB().Ping()
	migrate(db)
	return db
}
func createConnectionString(host string, port int, pwd, username, dbname string) string {
	return fmt.Sprintf("host=%s port=%v user=%s dbname=%s password=%s", host, port, username, dbname, pwd)
}
func migrate(db *gorm.DB) {
	db.AutoMigrate(Graph{}, Node{}, Edge{})
	db.Model(Node{}).AddForeignKey("graph_id", "graphs(id)", "RESTRICT", "RESTRICT")
	db.Model(Edge{}).AddForeignKey("from", "node(id)", "RESTRICT", "RESTRICT")
	db.Model(Edge{}).AddForeignKey("to", "node(id)", "RESTRICT", "RESTRICT")
}
