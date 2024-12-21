package store

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

type Store struct {
	Conn *pgx.Conn
}

func New(connString string) (*Store, error) {

	s := &Store{}

	if connString == "" {
		return nil, fmt.Errorf("database connection is not provided")
	}
	ctx := context.Background()

	db, err := pgx.Connect(ctx, connString)
	if err != nil {
		return nil, fmt.Errorf("Error is %v", err)
	}

	s.Conn = db
	fmt.Println("Connected to PostgreSQL!")
	return s, nil
}

func (s *Store) Close() {
	ctx := context.Background()
	if s.Conn != nil {
		_ = s.Conn.Close(ctx)
		fmt.Println("Database connection closed.")
	}
}
