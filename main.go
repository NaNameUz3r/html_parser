package main

import (
	"NaNameUz3r/sitemap_gen/links"
	"encoding/xml"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	targetUrl := flag.String("url", "https://wannahack.in/", "link for which you want to build a sitemap")
	searchDepth := flag.Int("depth", 10, "the maximum depth value of the link search traverse")
	outputFile := flag.String("out", "./site-map.xml", "output file path")
	flag.Parse()

	links := breadthFirstSearch(*targetUrl, *searchDepth)
	renderXml(links, *outputFile)

}

func renderXml(links []string, filepath string) {
	buildXml := urlset{
		Urls:  make([]loc, len(links)),
		Xmlns: xmlns,
	}

	for i, link := range links {
		buildXml.Urls[i] = loc{link}
	}

	file, err := os.Create(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprint(file, xml.Header)
	enc := xml.NewEncoder(file)
	enc.Indent("", "  ")
	if err := enc.Encode(buildXml); err != nil {
		panic(err)
	}
	fmt.Fprint(file, "\n")
}

type empty struct{}

func breadthFirstSearch(urlString string, maxDepth int) []string {
	seenMap := make(map[string]empty)
	var layer map[string]empty
	nextLayer := map[string]empty{
		urlString: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		layer, nextLayer = nextLayer, make(map[string]empty)
		if len(layer) == 0 {
			break
		}
		for url, _ := range layer {
			if _, ok := seenMap[url]; ok {
				continue
			}
			seenMap[url] = empty{}
			for _, link := range parseLinks(url) {
				if _, ok := seenMap[link]; !ok {
					nextLayer[link] = empty{}
				}
			}
		}
	}

	seenLinks := make([]string, 0, len(seenMap))
	for url, _ := range seenMap {
		seenLinks = append(seenLinks, url)
	}

	return seenLinks
}

func parseLinks(url string) []string {
	response, err := http.Get(url)
	if err != nil {
		return []string{}
	}
	defer response.Body.Close()

	baseUrl := getBaseUrl(response)
	allLinks, err := parseHrefs(response.Body, baseUrl)
	if err != nil {
		return []string{}
	}

	filteredLinks := filterUrls(allLinks, withPrefix(baseUrl))
	return filteredLinks
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
