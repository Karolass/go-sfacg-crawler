package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func GetHtml(URL string) (res *http.Response) {
	// sleep
	time.Sleep(RandomTime())

	client := &http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.94 Safari/537.36")

	res, err = client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return
}

func (catalogs *Catalogs) Get(URL string) (nextPage string) {

	res := GetHtml(URL)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatalln(err)
	}

	attr, exist := doc.Find("li.pagebarNext a").Attr("href")
	if exist {
		nextPage = attr
	}

	doc.
		Find(".Comic_Pic_List li:nth-child(2)").
		Each(func(i int, this *goquery.Selection) {

			title := this.Find("strong").Find("a").Text()
			attr, exist := this.Find("strong").Find("a").Attr("href")
			var ID string
			if exist {
				r := regexp.MustCompile(`^[\w\W]+\/HTML\/(\w+)\/$`)
				results := r.FindStringSubmatch(attr)
				if len(results) > 1 {
					ID = results[1]
				}
			}

			var author, category string
			this.
				Find("a.Blue_link1").
				Each(func(i int, this *goquery.Selection) {
					if i%2 == 0 {
						author = this.Text()
					} else {
						category = this.Text()
					}
				})

			description := strings.TrimSpace(this.Find("br").Get(2).NextSibling.Data)
			thumbnailURL, _ := this.Parent().Find("li.Conjunction").Find("img").Attr("src")

			if len(ID) > 0 {
				*catalogs = append(*catalogs, Catalog{
					ID:           ID,
					Title:        title,
					Author:       author,
					Category:     category,
					Description:  description,
					URL:          attr,
					ThumbnailURL: thumbnailURL,
				})
			}
		})

	return
}

func (chapters *Chapters) Get(catalogID string, URL string) {

	res := GetHtml(URL)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatalln(err)
	}

	doc.
		Find("ul.serialise_list li a").
		Each(func(i int, this *goquery.Selection) {

			var title, URL string
			if len(this.Text()) > 0 {
				title = this.Text()
			} else {
				title = this.Find("font").Text()
			}

			attr, exist := this.Attr("href")
			if exist {
				URL = "http://comic.sfacg.com" + attr
			}

			*chapters = append(*chapters, Chapter{
				CatalogID: catalogID,
				Title:     title,
				URL:       URL,
			})
		})

	return
}

func (pages *Pages) Get(catalogID string, chapterTitle string, URL string) {

	res := GetHtml(URL)
	doc, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		log.Fatalln(err)
	}

	attr, exist := doc.Find("script").Eq(1).Attr("src")
	if exist {
		jsURL := "http://comic.sfacg.com" + attr

		res = GetHtml(jsURL)
		bytes, _ := ioutil.ReadAll(res.Body)
		re := regexp.MustCompile(`\/Pic\/[\w|\/]+\.\w+`)
		for _, match := range re.FindAllString(string(bytes), -1) {
			*pages = append(*pages, Page{
				CatalogID:    catalogID,
				ChapterTitle: chapterTitle,
				URL:          "http://comic.sfacg.com" + match,
			})
		}
	}

	return
}
