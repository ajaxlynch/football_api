package database

import (
  "database/sql"
  _ "github.com/lib/pq"
  "fmt"
)

const (
  dbhost = DBHOST
  dbport = DBPORT
  dbuser = DBUSER
  dbpass = DBPASS
  dbname = DBNAME
)

var db *sql.DB

func InitDb() *sql.DB {
  // config := dbConfig()

  var err error
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
  dbhost, dbport, dbuser,  dbname)

  db, err = sql.Open("postgres", psqlInfo)
  if err != nil{
    panic(err)
  }

  fmt.Println(psqlInfo)
  fmt.Println("Successfully connected")

  return db
}
