/*
Main Entry Point for the scraper

Make sure images/original/ folder and config.json are placed with the binary.

Author: uTosTan
*/
package main

import (
	"log"
	"time"

	"bitbucket.org/utostan/mulukscraper/data"
	"bitbucket.org/utostan/mulukscraper/scrape"

	_ "github.com/go-sql-driver/mysql"
)

func run(s scrape.Scraper, categoryURL *string, c chan []data.News) {
	test := s.Scrape(categoryURL)
	c <- test
}

func initializeScrapers() map[string]scrape.Scraper {
	var sitemap data.SiteMap
	sitemap.Get()

	siteScrapers := make(map[string]scrape.Scraper)
	siteScrapers["BBC Nepal"] = &scrape.Bbc{Src: &data.Source{BaseURL: sitemap.Sites["BBC Nepal"]}}
	return siteScrapers
}

func main() {
	// Get CategorySite table from DB
	var css data.CategorySites
	css.Get()

	// Time
	start := time.Now()

	siteScrapers := initializeScrapers()

	var news []data.News

	// Iterate through the sites and initiate goroutines to scrape the URLs
	for _, cs := range css.CategorySites {
		if scraper, ok := siteScrapers[cs.Site.SiteName]; ok {
			c := make(chan []data.News)
			go run(scraper, &cs.CategoryURL, c)
			news = append(news, <-c...)
		}
	}

	elapsed := time.Since(start)

	log.Println("Scraping took: " + elapsed.String())
}
