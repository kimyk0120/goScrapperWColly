package main

import (
	"encoding/csv"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strconv"
	"time"
)

func Naver() {
	fName := "naver_news.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	_ = writer.Write([]string{"no", "head", "alink", "time"})

	// Instantiate default collector
	c := colly.NewCollector()

	//c.OnHTML(".news_box .sub_news .list_news", func(e *colly.HTMLElement) {
	c.OnHTML(".hdline_news", func(e *colly.HTMLElement) {
		headText := ""
		link := ""
		for i := 1; i <= 5; i++ {
			headText = e.ChildText("li:nth-child("+strconv.Itoa(i)+") .hdline_article_tit a")
			link = e.ChildAttr("li:nth-child("+strconv.Itoa(i)+") .hdline_article_tit a", "href")
			now := time.Now()
			println(headText)
			_ = writer.Write([]string{
				strconv.Itoa(i),
				headText,
				link,
				now.Format(time.ANSIC),
			})
		}
	})

	_ = c.Visit("https://news.naver.com/")
	log.Printf("Scraping finished, check file %q for results\n", fName)
}