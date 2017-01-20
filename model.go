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
	CatalogID string `json:"catalogID"`
	Title     string `json:"title"`
	URL       string
	ObjectId  string `json:"objectId"`
}

type Chapters []Chapter

type Page struct {
	CatalogID    string
	ChapterTitle string
	URL          string
}

type Pages []Page

type ChapterPage struct {
	Pages []string `json:"pages"`
}
