package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

type Job struct {
	Title   string `json:"title"`
	Company string `json:"company"`
}

func main() {
	// default allocator (առանց ExecPath)
	opts := append(chromedp.DefaultExecAllocatorOptions[:])

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// timeout
	ctx, cancel = context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	url := "https://staff.am/am/jobs"
	var jobs []Job

	log.Println("➡️ Navigating to:", url)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),

		// սպասում ենք, որ առնվազն 1 գործի վերնագիր երևա
		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("⌛ Waiting for job listings to load...")
			return nil
		}),
		chromedp.WaitVisible(`a[href*="/am/jobs/"]`, chromedp.ByQuery),

		chromedp.ActionFunc(func(ctx context.Context) error {
			log.Println("✅ Jobs are visible, extracting data...")
			return nil
		}),

		chromedp.EvaluateAsDevTools(`
			Array.from(document.querySelectorAll('div[data-id]')).map(item => {
				const title = item.querySelector('a[href*="/am/jobs/"] div')?.innerText.trim() || "";
				const company = item.querySelector('a[href*="/am/company/"] div')?.innerText.trim() || "";
				return { title, company };
			});
		`, &jobs),
	)
	if err != nil {
		log.Fatal("❌ Error during chromedp.Run:", err)
	}

	if len(jobs) == 0 {
		log.Println("⚠️ Warning: No jobs extracted, maybe selectors need update.")
	} else {
		log.Printf("✅ Extracted %d jobs\n", len(jobs))
	}

	data, _ := json.MarshalIndent(jobs, "", "  ")
	fmt.Println(string(data))
}
