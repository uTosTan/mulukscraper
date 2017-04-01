package scrape

import "bitbucket.org/utostan/mulukscraper/data"

type Scraper interface {
    Scrape() []data.News
    crawl(url string, ch chan *data.News, chDone chan bool)
}