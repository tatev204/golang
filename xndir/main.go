package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Title struct {
	Ru string `json:"ru"`
	En string `json:"en"`
	Am string `json:"am"`
}
type Company struct {
	Domain_id           int    `json:"domain_id"`
	Profile_image       string `json:"profile_image"`
	Id                  int    `json:"id"`
	Title               Title  `json:"title"`
	Is_unlimited        int    `json:"is_unlimited"`
	Is_verified         bool   `json:"is_verified"`
	Slug                string `json:"slug"`
	IsEmptyProfileImage bool   `json:"isEmptyProfileImage"`
}

type Category struct {
	Code  string `json:"code"`
	Id    int    `json:"id"`
	Title Title  `json:"title"`
}

type Job struct {
	Title           Title    `json:"title"`
	Deadline        string   `json:"deadline"`
	CompaniesStruct Company  `json:"companiesStruct"`
	Category        Category `json:"category"`
	Slug            Title    `json:"slug"`
	Status          int      `json:"status"`
	Saved_job_id    bool     `json:"saved_job_id"`
	Like_id         bool     `json:"like_id"`
	Is_featured     bool     `json:"is_featured"`
	LikeCount       int      `json:"likeCount"`
	Is_new          bool     `json:"is_new"`
	ItemType        int      `json:"itemType"`
	Is_following    bool     `json:"is_following"`
}

type PageProps struct {
	Jobs       []Job `json:"jobs"`
	TotalCount int   `json:"totalCount"`
}

type ApiResponse struct {
	PageProps PageProps `json:"pageProps"`
}

func fetchJobs() ([]Job, error) {
	const url = "https://staff.am/_next/data/QqNfKEDjEeO01hJlkXWg_/am/jobs.json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response ApiResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.PageProps.Jobs, nil
}

func main() {
	jobs, err := fetchJobs()
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Printf("Found %d jobs:\n", len(jobs))
	for i, job := range jobs {
		fmt.Printf("%d. %s - %s\n", i+1, job.Title.Am, job.CompaniesStruct.Title.Am)

	}
}
