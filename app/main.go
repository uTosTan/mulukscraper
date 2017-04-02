package main

import (
    "fmt"
    "time"
    "net/http"
    "bitbucket.org/utostan/mulukscraper/scrape"
    "bitbucket.org/utostan/mulukscraper/data"
)

func run(s scrape.Scraper, c chan []data.News) {
    test := s.Scrape("nepali/news")
    c <- test
}

func main() {
    var sites data.Sites
    sites.Get()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        start := time.Now()

        siteScrapers :=  make(map[string]scrape.Scraper)
        siteScrapers["BBC Nepal"] = &scrape.Bbc{&data.Source{"http://www.bbc.com/"}}
        
        var test []data.News

        for _, site := range sites.Sites {
            if scraper, ok := siteScrapers[site.SiteName]; ok {
                c := make(chan []data.News)
                go run(scraper, c)
                test = append(test, <-c...)
            }
        }

/*        for _, v := range siteScrapers {
            c := make(chan []data.News)
            go run(v, c)
            test = append(test, <-c...)
        }*/

        elapsed := time.Since(start)

        fmt.Println(elapsed)

        fmt.Fprint(w, test)
    })
    http.ListenAndServe(":8080", nil)
}