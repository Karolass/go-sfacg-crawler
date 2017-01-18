package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	// "log"
	// "time"
)

var catalogCount, chapterCount, pageCount = 0, 0, 0
var nowPage, limit = 1, 0

var (
	startFlag string
)

func init() {
	flag.StringVar(&startFlag, "start", "catalog", "start scrape catalog(Default), chapter or page")
	flag.Parse()
}

func main() {
	if startFlag == "catalog" {
		fmt.Printf("Scraping \"Catalog\"\n")
		run("http://comic.sfacg.com/Catalog/")
	}
	// start := time.Now()
	// run("http://comic.sfacg.com/Catalog/")
	// fmt.Printf("Total Catalog counts: %d \n", catalogCount)
	// fmt.Printf("Total Chapter counts: %d \n", chapterCount)
	// fmt.Printf("Total Page counts: %d \n", pageCount)
	// fmt.Println(time.Since(start))
}

func run(URL string) {
	catalogs := new(Catalogs)
	nextPage := catalogs.Get(URL)

	// bytes, err := json.MarshalIndent(catalogs, "", "    ")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// fmt.Println(string(bytes))
	// fmt.Println(nextPage)

	catalogCount += len(*catalogs)
	// get Chapters
	// runChatper(catalogs)

	for _, catalog := range *catalogs {
		catalog.create()
	}

	fmt.Printf("Scrape catalog page %d complete!\n", nowPage)
	if len(nextPage) > 0 && (limit == 0 || nowPage < limit) {
		nowPage++
		run(nextPage)
	}
}

func runChatper(catalogs *Catalogs) {
	chapters := new(Chapters)
	for _, catalog := range *catalogs {
		chapters.Get(catalog.ID, catalog.URL)

		// bytes, err := json.MarshalIndent(chapters, "", "    ")
		// if err != nil {
		//     log.Fatalln(err)
		// }

		// fmt.Println(catalog.Title)
		// fmt.Println(string(bytes))

		chapterCount += len(*chapters)
		// get Pages
		runPage(chapters)
	}
}

func runPage(chapters *Chapters) {
	pages := new(Pages)
	for _, chapter := range *chapters {
		pages.Get(chapter.CatalogID, chapter.Title, chapter.URL)

		// bytes, err := json.MarshalIndent(pages, "", "    ")
		// if err != nil {
		//     log.Fatalln(err)
		// }

		// fmt.Println(chapter.Title)
		// fmt.Println(string(bytes))

		pageCount += len(*pages)
	}
}
