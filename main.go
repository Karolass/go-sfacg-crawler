package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	// "log"
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
	switch startFlag {
	case "catalog":
		fmt.Printf("Scraping \"Catalog\"\n")
		run("http://comic.sfacg.com/Catalog/")
	case "chapter":
		fmt.Printf("Scraping \"Chapter\"\n")
		runChatper()
	case "test":
		fmt.Println(RandomTime())
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

	for _, catalog := range *catalogs {
		catalog.create()
	}

	fmt.Printf("Scrape catalog page %d complete!\n", nowPage)
	if len(nextPage) > 0 && (limit == 0 || nowPage < limit) {
		nowPage++
		run(nextPage)
	}
}

func runChatper() {
	catalogs := new(Catalogs)

	count := catalogs.count()
	var limit, skip = 100, 0
	for count > 0 {
		catalogs.find(skip, limit)
		for _, c := range *catalogs {
			chapters := new(Chapters)
			chapters.Get(c.ID, c.URL)
			for _, chapter := range *chapters {
				chapter.create()
			}
		}

		count -= limit
		skip += limit
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
