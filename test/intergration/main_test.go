package intergration

import (
	"database/sql"
	"golang/utils"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"golang/internal/model/domain"
)

var OpenConnection *domain.Queries
var TestDb *sql.DB

func TestMain(m *testing.M) {
	// load config to test

	config, err := utils.LoadConfig("../../")
	if err != nil {
		log.Fatal("Cannot load env  ", err)
	}

	TestDb, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot load connect db  ", err)
	}
	OpenConnection = domain.New(TestDb)
	os.Exit(m.Run())
}
