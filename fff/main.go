package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
)

func main() {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	resp, err := client.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	cookies := jar.Cookies(resp.Request.URL)
	fmt.Println("cookies")
	for _, c := range cookies {
		fmt.Printf("%s = %s\n", c.Name, c.Value)
	}

}
