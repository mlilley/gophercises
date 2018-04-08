package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/mlilley/gophercises/ex5"
)

func fetch(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

type SiteLink struct {
	Href  string
	Depth int
}

func mapPage(base url.URL, pageUrl string) []string {
	var pageLinks []SiteLink

	res, err := http.Get(pageLink.Href)
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	for _, l := range links {
		abs, err := base.Parse(l.Href)
		if err != nil {
			log.Fatal(err)
		}

		pageLinks = append(pageLinks, SiteLink{Href: (*abs).String(), Depth: d})
	}

	return pageLinks
}

func mapSite(siteURL string, depth int) ([]string, error) {
	base, err := url.Parse(siteURL)
	if err != nil {
		log.Fatal(err)
	}

	var sitemap []SiteLink

	var f func(string)
	f = func(pageURL string) {

		for _, pageUrl := range pageUrls {
			found := false
			for _, sitemapUrl := range sitemap {
				if pageUrl == sitemapUrl {
					found = true
					break
				}
			}
			if !found {
				sitemap = append(sitemap, pageUrl)
			}
		}

	}
	f(siteURL)

	return []string{""}, nil
}

func createSitemapXML([]string) string {
	return ""
}

func main() {
	siteUrl := flag.String("url", "", "Url of site to build sitemap from")
	flag.Parse()

	sitemap, err := mapSite(*siteUrl, 0)
	if err != nil {
		log.Fatal(err)
	}

	xml := createSitemapXML(sitemap)

	fmt.Println(xml)
}
