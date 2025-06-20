package main

import (
	"fmt"
	"strings"

	link "github.com/A-1evi/htmlLinkParser"
)

var exampleHtml = `<html>
  <body>
    <h1>Hello!</h1>
    <a href="/other-page">A link to another <span>page </span></a>
	 <a href="some-other-other-page">A link to another page</a>
  </body>
</html>
`

func main() {
	r := strings.NewReader(exampleHtml)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", links)
}
