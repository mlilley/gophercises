package link

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
	var links []Link

	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, parseLink(n))
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(node)

	return links, nil
}

func parseLink(node *html.Node) Link {
	return Link{Href: parseLinkHref(node), Text: parseLinkText(node)}
}

func parseLinkHref(n *html.Node) string {
	for _, attr := range n.Attr {
		if strings.ToLower(attr.Key) == "href" {
			return attr.Val
		}
	}
	return ""
}

func parseLinkText(n *html.Node) string {
	var txt = ""
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			txt += n.Data
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return strings.Join(strings.Fields(txt), " ")
}
