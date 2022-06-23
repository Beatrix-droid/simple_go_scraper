package main


import "fmt"
import "github.com/gocolly/colly"


func main(){

	//new collector initialised, wikipedia passed as base url
	collector := colly.NewCollector(colly.AllowedDomains("en.wikipedia.org"),)


	//find and print all links
	collector.OnHTML(".mw-parser-output", func(e*colly.HTMLElement){
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})

	collector.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}
