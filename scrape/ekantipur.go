package scrape

import (
    "bitbucket.org/utostan/mulukscraper/data"
)

type Ekantipur struct {
    Src *data.Source
}

func (ek *Ekantipur) Scrape() []data.News {
    var news []data.News

    news = append(news, data.News{ Headline: "Test2" })
    return news
}