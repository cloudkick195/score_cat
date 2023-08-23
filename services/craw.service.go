package services

import (
	"log"
	"score_cat/models"
	"score_cat/repositories"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type CrawService struct {
}

func (s *CrawService) CrawlAndSaveData() error {
	// Tạo kết nối tới trang web
	res, err := http.Get("https://bongdawap.com/")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Lặp qua các phần tử HTML chứa thông tin về đội bóng
	doc.Find("#lstmatch tr").Each(func(i int, s *goquery.Selection) {

		teamName := s.Find("#lstmatch .tendoi").Text()

		// Tạo mới team và lưu reference đến country và league
		team := models.Team{Name: teamName}
		repositories.DB.Create(&team)
	})

	return nil
}
