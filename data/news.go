package data

import (
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

type Category struct {
    CategoryId int
    CategoryName string
}

type CategorySite struct {
    Site Site
    Categorie Category
    CategoryUrl string
}

type Sites struct {
    Sites []Site
}

type CategorySites struct {
    CategorySites []CategorySite
}

type SiteMap struct {
    Sites map[string]string
}

func (sm *SiteMap) Get() {
    db := GetInstance().Connection
    rows, err := db.Query("SELECT SiteName, SiteUrl FROM Site WHERE IsActive=?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    sm.Sites = make(map[string]string)

    for rows.Next() {
        var name, url string
        if err := rows.Scan(&name, &url); err != nil {
            log.Fatal(err)
        }
        sm.Sites[name] = url
    }
}


func (n *Sites) Get() {
    db := GetInstance().Connection

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

func (cs *CategorySites) Get() {
    db := GetInstance().Connection

    rows, err := db.Query("SELECT Site.SiteId, Site.SiteUrl, Site.SiteName, Site.IsActive, CategorySite.CategoryUrl FROM Site INNER JOIN CategorySite ON CategorySite.SiteId=Site.SiteId WHERE Site.IsActive=?", 1)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    for rows.Next() {
        var i CategorySite
        if err := rows.Scan(&i.Site.SiteId, &i.Site.SiteUrl, &i.Site.SiteName, &i.Site.IsActive, &i.CategoryUrl); err != nil {
            log.Fatal(err)
        }
        cs.CategorySites = append(cs.CategorySites, i)
    }
}