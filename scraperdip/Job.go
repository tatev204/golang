package main

type Job struct {
	ID         int    `json:"id"`
	Deadline   string `json:"deadline"`
	IsRemote   bool   `json:"is_remote"`
	IsHot      bool   `json:"is_hot"`
	IsFeatured bool   `json:"is_featured"`
	Title      struct {
		En string `json:"en"`
		Am string `json:"am"`
	} `json:"title"`
	JobCity struct {
		Title struct {
			En string `json:"en"`
			Am string `json:"am"`
		} `json:"title"`
	} `json:"job_city"`
	CompaniesStruct struct {
		Title struct {
			En string `json:"en"`
			Am string `json:"am"`
		} `json:"title"`
		ProfileImage string `json:"profile_image"`
	} `json:"companiesStruct"`
}

type Response struct {
	PageProps struct {
		Jobs []Job `json:"jobs"`
	} `json:"pageProps"`
}
