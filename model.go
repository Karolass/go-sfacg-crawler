package main

type Catalog struct {
    ID           string
    Title        string
    Author       string
    Category     string
    Description  string
    URL          string
    ThumbnailURL string
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
