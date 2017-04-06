package data

import (
	"log"
)

// Driver interface for DB Models
type Driver interface {
	Get()
}

// News model
type News struct {
	Headline     string
	Excerpt      string
	NewsURL      string
	Body         string
	BodyOriginal string
	PublishDate  string
	Author       string
}

// Site model
type Site struct {
	SiteID   int
	SiteURL  string
	SiteName string
	IsActive bool
}

// Category model
type Category struct {
	CategoryID   int
	CategoryName string
}

// CategorySite model
type CategorySite struct {
	Site        Site
	Categorie   Category
	CategoryURL string
}

// Sites model (array of Site)
type Sites struct {
	Sites []Site
}

// CategorySites model (array of CategorySite)
type CategorySites struct {
	CategorySites []CategorySite
}

// SiteMap model (map of URLs)
type SiteMap struct {
	Sites map[string]string
}

// Get SiteMap
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

// Get Sites
func (n *Sites) Get() {
	db := GetInstance().Connection

	rows, err := db.Query("SELECT SiteId, SiteUrl, SiteName, IsActive FROM Site WHERE IsActive=?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var s Site
		if err := rows.Scan(&s.SiteID, &s.SiteURL, &s.SiteName, &s.IsActive); err != nil {
			log.Fatal(err)
		}
		n.Sites = append(n.Sites, s)
	}
}

// Get CategorySites
func (cs *CategorySites) Get() {
	db := GetInstance().Connection

	rows, err := db.Query("SELECT Site.SiteId, Site.SiteUrl, Site.SiteName, Site.IsActive, CategorySite.CategoryUrl FROM Site INNER JOIN CategorySite ON CategorySite.SiteId=Site.SiteId WHERE Site.IsActive=?", 1)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var i CategorySite
		if err := rows.Scan(&i.Site.SiteID, &i.Site.SiteURL, &i.Site.SiteName, &i.Site.IsActive, &i.CategoryURL); err != nil {
			log.Fatal(err)
		}
		cs.CategorySites = append(cs.CategorySites, i)
	}
}
