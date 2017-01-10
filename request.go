package main

import (
    "log"
    "net/http"
    "strings"
    "regexp"

    "github.com/PuerkitoBio/goquery"
)

func GetCatalogs(URL string) (catalogs Catalogs, nextPage string){
    client := &http.Client{}

    req, err := http.NewRequest("GET", URL, nil)
    if err != nil {
        log.Fatalln(err)
    }

    req.Header.Set("User-Agent",
        "Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/40.0.2214.94 Safari/537.36")

    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    doc, err := goquery.NewDocumentFromReader(resp.Body)
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
                if i % 2 == 0 {
                    author = this.Text()
                } else {
                    category = this.Text()
                }
            })

        description := strings.TrimSpace(this.Find("br").Get(2).NextSibling.Data)
        thumbnailURL, _ := this.Parent().Find("li.Conjunction").Find("img").Attr("src")

        if len(ID) > 0 {
            catalogs = append(catalogs, Catalog{
                        ID: ID,
                        Title: title,
                        Author: author,
                        Category: category,
                        Description: description,
                        URL: attr,
                        ThumbnailURL: thumbnailURL,
                    })
        }
    })

    return
}