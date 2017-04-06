package scrape

import (
	//"fmt"
	"strconv"
	"strings"
	"time"

	"bitbucket.org/utostan/mulukscraper/data"
	"bitbucket.org/utostan/mulukscraper/image"
	"github.com/PuerkitoBio/goquery"
)

// Bbc struct
type Bbc struct {
	Src *data.Source // Is this really necessary..
}

func (bbc *Bbc) crawl(url string, ch chan *data.News, chDone chan bool) {
	defer func() {
		chDone <- true
	}()

	var tm time.Time
	doc, _ := goquery.NewDocument(url)
	headline := doc.Find("div.story-body h1.story-body__h1").Text()
	timestamp, ok := doc.Find("div.date.date--v2").Attr("data-seconds")
	storyBody := doc.Find("div.story-body__inner")

	if ok {
		i, _ := strconv.ParseInt(timestamp, 10, 64)
		tm = time.Unix(i, 0)
	} else {
		tm = time.Now()
	}

	imageURL, ok := storyBody.Find("figure img").Attr("src")

	storyBody.Find("div").Remove()
	body, _ := storyBody.Html()

	if ok {
		s := strings.Split(imageURL, "/")
		image.Get(&imageURL, "images/"+s[len(s)-1])
	}

	ch <- &data.News{Headline: headline, BodyOriginal: body, PublishDate: tm.String()}
}

func skip(chDone chan bool) {
	chDone <- true
}

// Scrape Bbc
func (bbc *Bbc) Scrape(category *string) []data.News {
	doc, _ := goquery.NewDocument(bbc.Src.BaseURL + *category)
	chNews := make(chan *data.News)
	chDone := make(chan bool)
	var news []data.News

	blocks := doc.Find("div.eagle div.eagle-item.faux-block-link").Each(func(i int, s *goquery.Selection) {
		a := s.Find("div.eagle-item__body").Find("a.title-link")

		link, ok := a.Attr("href")

		if ok && a.Find("span.off-screen").Length() < 1 {
			go bbc.crawl(bbc.Src.BaseURL+link, chNews, chDone)
		} else {
			go skip(chDone)
		}
	})

	for i := 0; i < blocks.Length(); {
		select {
		case n := <-chNews:
			news = append(news, *n)
		case <-chDone:
			i++
		}
	}

	return news
}
