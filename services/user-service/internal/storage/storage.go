package storage

import (
	"database/sql"

	"context"

	"fmt"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type StorageIF interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, user *User) (*User, error)
	DeleteUser(ctx context.Context, user *User) (*User, error)
}

type Store struct {
	db  *sql.DB
	log *logrus.Logger
}

func NewDB(host, port, user, password, database string) *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func NewStore(db *sql.DB, log *logrus.Logger) *Store {
	s := Store{db: db, log: log}
	return &s
}

func (s *Store) CreateUser(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (s *Store) GetUser(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (s *Store) UpdateUser(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}

func (s *Store) DeleteUser(ctx context.Context, user *User) (*User, error) {
	return nil, nil
}
