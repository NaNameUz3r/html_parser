package main

import (
	"fmt"
	links "links_parser"
	"strings"
)

var testHtml = `
<html>
<body>
	<h3>Hello world</h3>
	<a href="https://shorturl.at/eFKY2">Just a link
		<p> to some strange </p>
		but beautifull and <b> COMPLETELY NORMAL </b>
	site
	</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(testHtml)
	links, err := links.Parse(r)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)
}
