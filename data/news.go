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


func (n *Sites) Get() {
    db := GetInstance().Connection;

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
    db := GetInstance().Connection;

    rows, err := db.Query("SELECT site.SiteId, site.SiteUrl, site.SiteName, site.IsActive, categorysite.CategoryUrl FROM site INNER JOIN categorysite ON categorysite.SiteId=site.SiteId WHERE site.IsActive=?", 1)
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