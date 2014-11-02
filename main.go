package main

import (
	"os"
	"flag"
	"strings"
	"net/url"
	"github.com/skratchdot/open-golang/open"
)

func main() {
	args := os.Args[1:]
	w := strings.Join(args, " ")
	t := flag.Int("t", 0, "Search type (0:normal, 1:image)")
	flag.Parse()

	var href *url.URL
	href, err := url.Parse("http://www.google.co.jp")
	if err != nil {
		panic("Ouch!")
	}
	href.Path += "/search"

	params := url.Values{}
	params.Add("q", w)
	if *t == 1 {
		params.Add("tbm", "isch")
	}

	href.RawQuery = params.Encode()
	open.Run(href.String())
}
