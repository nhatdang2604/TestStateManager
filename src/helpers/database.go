package helpers

import (
	"database/sql"
	"log"
	"sync"
	"time"

	"github.com/nhatdang2604/TestStateManager/src/constants"
)

type DB struct {
	ConnectionString string
	Connection       *sql.DB
}

//Singleton db connection
var db *DB
var once sync.Once

func NewDB(
	connectionString string,
	maxConnectionLifeTimeInMinute int,
	maxOpenConnection int,
	maxIdleConnection int) *DB {

	db := &DB{
		ConnectionString: connectionString,
	}

	conn, err := sql.Open("mysql", db.ConnectionString)
	if nil != err {
		log.Fatalf("Error: error on opening db connection: %v", err)
		return nil
	}

	db.Connection = conn
	db.Connection.SetConnMaxLifetime(time.Minute * time.Duration(maxConnectionLifeTimeInMinute))
	db.Connection.SetMaxOpenConns(maxOpenConnection)
	db.Connection.SetMaxIdleConns(maxIdleConnection)

	return db
}

func InitDB() *DB {
	var db *DB = nil
	config := constants.GetConfigConstant()
	if nil != config {
		log.Fatalf("Error: error on loading config constant, from database helper")
		return db
	}

	var user string = config.DBUser
	var pass string = config.DBPass
	var protocol string = config.DBProtocol
	var address string = config.DBHost + ":" + config.DBPort
	var dbName string = config.DBName

	connectionString := user + ":" + pass + "@" + protocol + "(" + address + ")/" + dbName + "?charset=utf8"
	db = NewDB(
		connectionString,
		config.DBMaxConnectionLifeTimeInMinute,
		config.DBMaxOpenConnection,
		config.DBMaxIdleConnection)

	return db
}

func GetDB() *DB {
	once.Do(func() {
		db = InitDB()
	})

	return db
}
