/*
Main Entry Point for the scraper

Make sure images/original/ folder and config.json are placed with the binary.

Author: uTosTan

 */
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
    // Get CategorySite table from DB
    var css data.CategorySites 
    css.Get()

    // Time
    start := time.Now()

    // Build a map of scrapers
    // They will be used based on the IsActive flag (in Sites)
    // Get BaseUrl from DB or Config (if possible)
    // SiteName will be the Unique Identifier
    siteScrapers :=  make(map[string]scrape.Scraper)
    siteScrapers["BBC Nepal"] = &scrape.Bbc{&data.Source{"http://www.bbc.com/"}}
    
    var test []data.News

    // Iterate through the sites and initiate goroutines to scrape the URLs
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