package data

import (
    "database/sql"
    _"github.com/go-sql-driver/mysql"
    "log"
)

type Driver interface {
    Get()
}

type News struct {
    Headline string
    Excerpt string
    NewsUrl string
    Body string
    BodyOriginal string
    PublishDate string
    Author string
}

type Site struct {
    SiteId int
    SiteUrl string
    SiteName string
    IsActive bool
}

type Sites struct {
    Sites []Site
}

func (n *Sites) Get() {
    db, err := sql.Open("mysql", "scrapenepal:suraj!2@/scrapenepal")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    rows, err := db.Query("SELECT SiteId, SiteUrl, SiteName, IsActive FROM Site WHERE IsActive=?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var s Site
        if err := rows.Scan(&s.SiteId, &s.SiteUrl, &s.SiteName, &s.IsActive); err != nil {
            log.Fatal(err)
        }
        n.Sites = append(n.Sites, s)
    }
}