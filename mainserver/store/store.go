package store

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct{
	Conn *sqlx.DB
}


func New(connString string)(*Store,error ){
	
	s :=&Store{}

	if(connString == ""){
		return nil,fmt.Errorf("database connection is not provided")
	}

	db, err := sqlx.Connect("postgres", connString)
	if err != nil {
		return nil,fmt.Errorf("Error is %v",err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	s.Conn = db
	fmt.Println("Connected to PostgreSQL!")
	return s,nil
}

func (s *Store) Close() {
	if s.Conn != nil {
		_ = s.Conn.Close()
		fmt.Println("Database connection closed.")
	}
}