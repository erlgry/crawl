package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
)

var (
	target, output string
	depth          int
)

func init() {
	const (
		targetUsage   = "A valid HTTP(s) URL to be crawled required)"
		depthDefault  = 5
		depthUsage    = "Max number of levels to crawl. Use 0 for no depth"
		outputDefault = "text"
		outputUsage   = "Output format (text or json)"
	)
	flag.StringVar(&target, "target", "", targetUsage)
	flag.StringVar(&target, "t", "", targetUsage+" (shorthand)")
	flag.IntVar(&depth, "depth", depthDefault, depthUsage)
	flag.IntVar(&depth, "d", depthDefault, depthUsage+" (shorthand)")
	flag.StringVar(&output, "output", outputDefault, outputUsage)
	flag.StringVar(&output, "o", outputDefault, outputUsage+" (shorthand)")
}

func main() {
	// parse the args
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	flag.Parse()

	// validate args
	invalid := func(message string) {
		fmt.Println(message)
		flag.Usage()
	}
	if target == "" {
		invalid("Missing target. A valid HTTP(s) URL must be specified")
	}
	t, err := url.Parse(target)
	if err != nil {
		invalid(fmt.Sprintf("Invalid target %s. Target must be a valid URL. %s", target, err))
	}
	if !t.IsAbs() {
		invalid(fmt.Sprintf("Invalid target %s. Target must be an absolute URL", target))
	}

	if depth < 0 {
		invalid(fmt.Sprintf("Invalid depth %d. Must be 0 or more", depth))
	}

	if output != "text" && output != "json" {
		invalid(fmt.Sprintf("Invalid output format %s. Must be 'text' or 'json'", output))
	}

	// TODO: kick off the crawler with root URL
}
