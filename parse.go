package link

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Links represent a Link  (<a href=".."> ) from the html deocument

type Link struct {
	href string
	text string
}

// html.Parse return html nodes
// Parsing is done by calling Parse with an io.Reader, which returns the root of the parse tree (the document element) as a *Node. It is the caller's responsibility to ensure that the Reader provides UTF-8 encoded HTML. For example, to process each anchor node in depth-first order:
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	nodes := linkNodes(doc)
	var links []Link
	for _, node := range nodes {
		links = append(links, BuildLink(node))

	}

	return links, nil
}

func BuildLink(n *html.Node) Link {
	var res Link
	for _, attr := range n.Attr {
		if attr.Key == "href" {
			res.href = attr.Val
			break
		}
	}
	res.text = text(n)
	return res
}

func text(n *html.Node) string {
	var res string
	if n.Type == html.TextNode {
		res = n.Data
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res += text(c) + " "
	}
	return strings.Join(strings.Fields(res), " ")
}

func linkNodes(n *html.Node) []*html.Node {
	var res []*html.Node
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, linkNodes(c)...)

	}
	return res
}
