package scrape

import (
    "bitbucket.org/utostan/mulukscraper/data"
    "github.com/PuerkitoBio/goquery"
    "time"
    "strconv"
)

type Bbc struct {
    Src *data.Source
}

func crawl(url string, ch chan *data.News, chDone chan bool) {
    defer func() {
        chDone <- true
    }()

    var tm time.Time
    doc, _ := goquery.NewDocument(url)

    headline := doc.Find("div.story-body h1.story-body__h1").Text()
    timestamp, ok := doc.Find("div.date.date--v2").Attr("data-seconds")

    if ok {
        i, _ := strconv.ParseInt(timestamp, 10, 64)
        tm = time.Unix(i, 0)
    } else {
        tm = time.Now()
    }

    doc.Find("div.story-body__inner").Find("div").Remove()
    body, _ := doc.Find("div.story-body__inner").Html()

    ch <- &data.News{Headline: headline, BodyOriginal: body, PublishDate: tm.String()}
}

func skip(chDone chan bool) {
    chDone <- true
}

func (bbc *Bbc) Scrape() []data.News {
    doc, _ := goquery.NewDocument(bbc.Src.BaseUrl + bbc.Src.Category)
    chNews := make(chan *data.News)
    chDone := make(chan bool)
    var news []data.News

    blocks := doc.Find("div.eagle div.eagle-item.faux-block-link").Each(func(i int, s *goquery.Selection) {
        a := s.Find("div.eagle-item__body").Find("a.title-link")

        link, ok := a.Attr("href")

        if ok && a.Find("span.off-screen").Length() < 1 {
            go crawl(bbc.Src.BaseUrl + link, chNews, chDone)
        } else {
            go skip(chDone)
        }
    })

    for i := 0; i < blocks.Length(); {
        select {
        case n := <-chNews:
            news = append(news, *n)
        case <-chDone:
            i++
        }
    }

    return news
}