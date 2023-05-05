package main

import (
	"NaNameUz3r/sitemap_gen/links"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	targetUrl := flag.String("url", "https://wannahack.in/", "link for which you want to build a sitemap")
	flag.Parse()

	response, err := http.Get(*targetUrl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	baseUrl := getBaseUrl(response)
	allLinks, err := parseHrefs(response.Body, baseUrl)
	if err != nil {
		panic(err)
	}

	filteredPages := filterUrls(allLinks, withPrefix(baseUrl))
	for _, p := range filteredPages {
		fmt.Println(p)
	}
}

func getBaseUrl(r *http.Response) string {
	requestUrl := r.Request.URL
	baseUrl := &url.URL{
		Scheme: requestUrl.Scheme,
		Host:   requestUrl.Host,
	}
	base := baseUrl.String()
	return base
}

func parseHrefs(body io.Reader, baseUrl string) ([]string, error) {
	links, err := links.Parse(body)
	if err != nil {
		return nil, err
	}
	var parsedLinks []string
	for _, link := range links {
		switch {
		case strings.HasPrefix(link.Href, "/"):
			parsedLinks = append(parsedLinks, baseUrl+link.Href)
		case strings.HasPrefix(link.Href, "http"):
			parsedLinks = append(parsedLinks, link.Href)
		default:
		}
	}
	return parsedLinks, nil
}

func filterUrls(links []string, keepFilter func(string) bool) []string {
	var filteredLinks []string

	for _, link := range links {
		if keepFilter(link) {
			filteredLinks = append(filteredLinks, link)
		}
	}
	return filteredLinks

}

func withPrefix(p string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, p)
	}
}
