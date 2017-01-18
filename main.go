package main

import (
	"fmt"
	// "encoding/json"
	// "log"
	"time"
)

var catalogCount, chapterCount, pageCount = 0, 0, 0
var nowPage, limit = 1, 5

func main() {
	start := time.Now()
	run("http://comic.sfacg.com/Catalog/")
	fmt.Printf("Total Catalog counts: %d \n", catalogCount)
	fmt.Printf("Total Chapter counts: %d \n", chapterCount)
	fmt.Printf("Total Page counts: %d \n", pageCount)
	fmt.Println(time.Since(start))
}

func run(URL string) {
	catalogs, nextPage := GetCatalogs(URL)

	// bytes, err := json.MarshalIndent(catalogs, "", "    ")
	// if err != nil {
	//     log.Fatalln(err)
	// }

	// fmt.Println(string(bytes))
	// fmt.Println(nextPage)

	catalogCount += len(catalogs)
	// get Chapters
	runChatper(catalogs)

	fmt.Printf("Scrape catalog page %d complete!\n", nowPage)
	if len(nextPage) > 0 && nowPage < limit {
		nowPage++
		run(nextPage)
	}
}

func runChatper(catalogs Catalogs) {
	for _, catalog := range catalogs {
		chapters := GetChapters(catalog.ID, catalog.URL)

		// bytes, err := json.MarshalIndent(chapters, "", "    ")
		// if err != nil {
		//     log.Fatalln(err)
		// }

		// fmt.Println(catalog.Title)
		// fmt.Println(string(bytes))

		chapterCount += len(chapters)
		// get Pages
		runPage(chapters)
	}
}

func runPage(chapters Chapters) {
	for _, chapter := range chapters {
		pages := GetPages(chapter.CatalogID, chapter.Title, chapter.URL)

		// bytes, err := json.MarshalIndent(pages, "", "    ")
		// if err != nil {
		//     log.Fatalln(err)
		// }

		// fmt.Println(chapter.Title)
		// fmt.Println(string(bytes))

		pageCount += len(pages)
	}
}
