package main

import (
	"fmt"
	"math"
	"net/url"
	"testing"

	"github.com/gonum/graph/encoding/dot"
	"github.com/gonum/graph/simple"
)

var nodes = []Node{
	{id: 0, url: toURL("/")},
	{id: 1, url: toURL("/1")},
	{id: 2, url: toURL("/2")},
	{id: 3, url: toURL("/3")},
	{id: 4, url: toURL("/4")},
	{id: 5, url: toURL("/5")},
	{id: 6, url: toURL("/6")},
	{id: 7, url: toURL("/7")},
	{id: 8, url: toURL("/8")},
	{id: 9, url: toURL("/9")},
}

var links = [][]int{
	{1, 2, 3},
	{0, 4, 7},
	{5},
	{6},
	{8},
	{9},
}

func TestPrintGraph(t *testing.T) {
	dg := simple.NewDirectedGraph(0, math.Inf(1))
	edges := edges()
	for _, e := range edges {
		dg.SetEdge(e)
	}
	got, _ := dot.Marshal(dg, "sitemap", "", "\t", true)
	fmt.Printf("%s", got)
}

func toURL(path string) *url.URL {
	return &url.URL{Scheme: "http", Host: "x.com", Path: path}
}

func edges() []Edge {
	edges := []Edge{}
	for f, ts := range links {
		for _, t := range ts {
			e := Edge{
				F: nodes[f],
				T: nodes[t],
			}
			edges = append(edges, e)
		}
	}
	return edges
}
