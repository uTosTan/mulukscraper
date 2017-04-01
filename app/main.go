package main

import (
    "fmt"
    "time"
    "net/http"
    "bitbucket.org/utostan/mulukscraper/scrape"
    "bitbucket.org/utostan/mulukscraper/data"
)

func run(s scrape.Scraper, c chan []data.News) {
    test := s.Scrape()
    c <- test
}

func main() {
    

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

        start := time.Now()

        sites :=  make(map[string]scrape.Scraper)
        sites["bbc"] = &scrape.Bbc{&data.Source{"http://www.bbc.com", "/nepali/news"}}
        sites["ekantipur"] = &scrape.Ekantipur{&data.Source{"bye", "bye2"}}
        var test []data.News

        for _, v := range sites {
            c := make(chan []data.News)
            go run(v, c)
            test = append(test, <-c...)
        }

        elapsed := time.Since(start)

        fmt.Println(elapsed)
        
        fmt.Fprint(w, test)
    })
    http.ListenAndServe(":8080", nil)
}