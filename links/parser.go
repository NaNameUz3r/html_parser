package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(r io.Reader) ([]Link, error) {
	document, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := linkNodes(document)
	links := make([]Link, 0)
	for _, node := range nodes {
		links = append(links, fetchLink(node))
	}
	return links, nil
}

func fetchLink(n *html.Node) Link {
	var link Link
	for _, attribute := range n.Attr {
		if attribute.Key == "href" {
			link.Href = attribute.Val
			break
		}
	}
	link.Text = parseLinkText(n)
	return link
}

func parseLinkText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var linkText string
	for child := n.FirstChild; child != nil; child = child.NextSibling {
		linkText += parseLinkText(child)
	}
	return normalizeText(linkText)
}

func normalizeText(s string) string {
	wordSlice := strings.Fields(s)
	return strings.Join(wordSlice, " ")
}

func linkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var linkNodesList []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodesList = append(linkNodesList, linkNodes(c)...)
	}

	return linkNodesList
}
