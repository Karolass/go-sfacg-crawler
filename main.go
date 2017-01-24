package main

import (
	"flag"
	"fmt"
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
		fmt.Println("Scraping \"Catalog\"")
		run("http://comic.sfacg.com/Catalog/")
	case "chapter":
		fmt.Println("Scraping \"Chapter\"")
		runChatper()
	case "page":
		fmt.Println("Scraping \"Page\"")
		runPage()
	case "pagenull":
		fmt.Println("Scraping \"Page for Chapter pages null\"")
		runPageNull()
	case "test":
		pages := new(Pages)
		pages.Get("", "", "http://comic.sfacg.com/HTML/juxingbs/023j/")
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
			fmt.Printf("Scrape \"%s\" chapters complete! Creating Parse data\n", c.Title)
			for _, chapter := range *chapters {
				objId := chapter.create()
				chapter.addRelation(objId, c.ObjectId)
			}
		}

		count -= limit
		skip += limit
	}
}

func runPage() {
	chapters := new(Chapters)

	count := chapters.count(false)
	var limit, skip = 100, 0
	for count > 0 {
		chapters.find(skip, limit, false)
		for _, c := range *chapters {
			pages := new(Pages)
			pages.Get(c.CatalogID, c.Title, c.URL)
			if len(*pages) > 0 {
				fmt.Printf("Scrape \"%s\" chapter \"%s\" pages complete! Creating Parse data\n", c.CatalogID, c.Title)

				chapterPage := new(ChapterPage)
				for _, page := range *pages {
					chapterPage.Pages = append(chapterPage.Pages, page.URL)
				}
				c.update(*chapterPage)
			} else {
				fmt.Printf("Scrape \"%s\" chapter \"%s\" no page data\n", c.CatalogID, c.Title)
			}
		}

		count -= limit
		skip += limit
	}
}

func runPageNull() {
	chapters := new(Chapters)

	count := chapters.count(true)
	var limit, skip = 100, 0
	for count > 0 {
		chapters.find(skip, limit, true)
		for _, c := range *chapters {
			pages := new(Pages)
			pages.Get(c.CatalogID, c.Title, c.URL)

			if len(*pages) > 0 {
				fmt.Printf("Scrape \"%s\" chapter \"%s\" pages complete! Creating Parse data\n", c.CatalogID, c.Title)

				chapterPage := new(ChapterPage)
				for _, page := range *pages {
					chapterPage.Pages = append(chapterPage.Pages, page.URL)
				}
				c.update(*chapterPage)
			} else {
				fmt.Printf("Scrape \"%s\" chapter \"%s\" no page data\n", c.CatalogID, c.Title)
			}
		}

		count -= limit
		skip += limit
	}
}
