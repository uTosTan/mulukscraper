package main

import (
    "log"
    "time"
    "bitbucket.org/utostan/mulukscraper/scrape"
    "bitbucket.org/utostan/mulukscraper/data"
)

func run(s scrape.Scraper, categoryUrl *string, c chan []data.News) {
    test := s.Scrape(categoryUrl)
    c <- test
}

func main() {
    var css data.CategorySites
    css.Get()

    // Time
    start := time.Now()

    siteScrapers :=  make(map[string]scrape.Scraper)
    siteScrapers["BBC Nepal"] = &scrape.Bbc{&data.Source{"http://www.bbc.com/"}}
    
    var test []data.News

    for _, cs := range css.CategorySites {
        if scraper, ok := siteScrapers[cs.Site.SiteName]; ok {
            c := make(chan []data.News)
            go run(scraper, &cs.CategoryUrl, c)
            test = append(test, <-c...)
        }
    }

    elapsed := time.Since(start)

    log.Println("Scraping took: " + elapsed.String())
}