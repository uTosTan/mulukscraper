package scrape

import "bitbucket.org/utostan/mulukscraper/data"

type Scraper interface {
    Scrape() []data.News
}