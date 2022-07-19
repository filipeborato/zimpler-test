package service

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GoQueryCandy() error {
	webPage := "https://candystore.zimpler.net/"
	resp, err := http.Get(webPage)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return err
	}

	title := doc.Find("title").Text()
	fmt.Println(title)

	return nil
}
