package main

type Catalog struct {
    ID string
    Title string
    Author string
    Category string
    Description string
    URL string
    ThumbnailURL string
}

type Catalogs []Catalog
