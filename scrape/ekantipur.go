package scrape

import (
	"bitbucket.org/utostan/mulukscraper/data"
)

// Ekantipur struct
type Ekantipur struct {
	Src *data.Source
}

func (ek *Ekantipur) crawl(url string, ch chan *data.News, chDone chan bool) {

}

// Scrape Ekantipur
func (ek *Ekantipur) Scrape(category *string) []data.News {
	var news []data.News

	news = append(news, data.News{Headline: "Test2"})
	return news
}
