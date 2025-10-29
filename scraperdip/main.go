package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"io"
	"net/http"
	"os"
)

func main() {
	page := 3
	lang := "am"
	url := fmt.Sprintf("https://staff.am/_next/data/k2-yhN689sCtAyL-xiMPd/%s/jobs.json?page=%d", lang, page)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("HTTP request error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	var data Response
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	conn, err := pgx.Connect(context.Background(),
		"postgres://postgres:tatev1234@localhost:5432/JobsDB?sslmode=disable")
	if err != nil {
		fmt.Println("Database connection error:", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	for _, job := range data.PageProps.Jobs {
		_, err := conn.Exec(context.Background(), `
			INSERT INTO jobs 
				(job_id, title_en, title_am, city_en, city_am, company_en, company_am, company_image, is_remote, deadline)
			VALUES 
				($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
			ON CONFLICT (job_id) DO NOTHING;
		`,
			job.ID,
			job.Title.En, job.Title.Am,
			job.JobCity.Title.En, job.JobCity.Title.Am,
			job.CompaniesStruct.Title.En, job.CompaniesStruct.Title.Am,
			job.CompaniesStruct.ProfileImage,
			job.IsRemote,
			job.Deadline,
		)
		if err != nil {
			fmt.Println("Insert error:", err)
			continue
		}
		fmt.Printf("✅ Added job: %s at %s\n", job.Title.En, job.CompaniesStruct.Title.En)
	}

	fmt.Println("All jobs inserted successfully!")
}
