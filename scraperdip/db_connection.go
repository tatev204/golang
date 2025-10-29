package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

func db_conn() {
	// change username, password, dbname if needed
	url := "postgres://postgres:tatev1234@localhost:5432/JobsDB?sslmode=disable"

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatal("❌ Cannot connect:", err)
	}
	defer conn.Close(context.Background())

	fmt.Println("✅ Connected to PostgreSQL!")

	// test query
	var message string
	err = conn.QueryRow(context.Background(), "SELECT 'Hello from PostgreSQL!'").Scan(&message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)

	// 1 request -> response  List<JobsResponse>
	// 2  List forEach, for, -> _, err = conn.Exec(context.Background(),
	//		"INSERT INTO jobs (title) VALUES ($1)",
	//		"Android Developer")

	_, err = conn.Exec(context.Background(),
		"INSERT INTO jobs (title) VALUES ($1)",
		"Android Developer")
	if err != nil {
		log.Fatal("❌ Insert failed:", err)
	}

	fmt.Println("✅ Test row inserted!")
}

type JobResponse struct {
	jobTitle    string
	salary      int
	companyName string
}
