package main

import (
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_migrations")
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db

	// migrate create -ext sql -dir db/migrations create_table_first
	// migrate create -ext sql -dir db/migrations create_table_second
	// migrate create -ext sql -dir db/migrations create_table_third
	// migrate create -ext sql -dir db/migrations sample_dirty_state
	// migrate -database "mysql://root@tcp(localhost:3306)/go_migrations" -path db/migrations up
	// migrate -database "mysql://root@tcp(localhost:3306)/go_migrations" -path db/migrations down
	// migrate -database "mysql://root@tcp(localhost:3306)/go_migrations" -path db/migrations version
	// migrate -database "mysql://root@tcp(localhost:3306)/go_migrations" -path db/migrations force 20220922043738
}
