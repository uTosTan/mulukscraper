package scrape

import "bitbucket.org/utostan/mulukscraper/data"

type Scraper interface {
    Scrape(category *string) []data.News
    crawl(url string, ch chan *data.News, chDone chan bool)
}