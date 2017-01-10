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

    if len(nextPage) > 0 {
        run(nextPage)
    }
}