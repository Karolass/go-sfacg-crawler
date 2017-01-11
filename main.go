package main

import (
    "fmt"
    "encoding/json"
    "log"
    "time"
)

func main() {
    start := time.Now()
    run("http://comic.sfacg.com/Catalog/")
    fmt.Println(time.Since(start))
}

func run(URL string) {
    catalogs, nextPage := GetCatalogs(URL)

    bytes, err := json.MarshalIndent(catalogs, "", "    ")
    if err != nil {
        log.Fatalln(err)
    }

    fmt.Println(string(bytes))
    fmt.Println(nextPage)

    // get Chapters
    runChatper(catalogs)

    // if len(nextPage) > 0 {
    //     run(nextPage)
    // }
}

func runChatper(catalogs Catalogs) {
    for _, catalog := range catalogs {
        chapters := GetChapters(catalog.ID, catalog.URL)

        bytes, err := json.MarshalIndent(chapters, "", "    ")
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println(catalog.Title)
        fmt.Println(string(bytes))

        // get Pages
        runPage(chapters)
    }
}

func runPage(chapters Chapters) {
    for _, chapter := range chapters {
        pages := GetPages(chapter.CatalogID, chapter.Title, chapter.URL)

        bytes, err := json.MarshalIndent(pages, "", "    ")
        if err != nil {
            log.Fatalln(err)
        }

        fmt.Println(chapter.Title)
        fmt.Println(string(bytes))
    }
}
