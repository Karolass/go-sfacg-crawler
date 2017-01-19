package main

type Catalog struct {
	ID           string
	Title        string `json:"title"`
	Author       string `json:"author"`
	Category     string `json:"category"`
	Description  string `json:"description"`
	URL          string
	ThumbnailURL string `json:"thumbnailURL"`
	ObjectId     string `json:"objectId"`
}

type Catalogs []Catalog

type Chapter struct {
	CatalogID string
	Title     string
	URL       string
}

type Chapters []Chapter

type Page struct {
	CatalogID    string
	ChapterTitle string
	URL          string
}

type Pages []Page
