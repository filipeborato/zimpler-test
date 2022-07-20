package service

import (
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"zimpler-test/entity"

	"github.com/PuerkitoBio/goquery"
)

func GoQueryCandy() (entity.TopRate, error) {
	topRate := entity.TopRate{}
	webPage := "https://candystore.zimpler.net/"
	resp, err := http.Get(webPage)

	if err != nil {
		return entity.TopRate{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("failed to fetch data: ", resp.StatusCode, resp.Status)
		return entity.TopRate{}, err
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return entity.TopRate{}, err
	}

	doc.Find("table.summary").Find("tbody tr").Each(func(i int, s *goquery.Selection) {
		str, err := s.Html()
		if err != nil {
			log.Println(err)
			return
		}

		split := strings.Split(str, "\"")
		totalSnacks, err := strconv.Atoi(split[1])
		if err != nil {
			log.Println(err)
			return
		}

		text := s.Text()
		text = strings.ReplaceAll(text, " ", "")
		textSplit := strings.Split(text, "\n")

		resp := entity.StoreResp{}
		resp.Name = textSplit[1]
		resp.FavouriteSnack = textSplit[2]
		resp.TotalSnacks = totalSnacks

		topRate = append(topRate, resp)
	})

	sort.SliceStable(topRate, func(i, j int) bool {
		return topRate[i].TotalSnacks > topRate[j].TotalSnacks
	})
	log.Println("####### topRate = ", topRate)

	return topRate, nil
}
