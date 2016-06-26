package main

import (
	"net/url"

	"github.com/gonum/graph"
)

// Node represents a node in the sitemap
type Node struct {
	id    int
	url   *url.URL
	html  bool
	title string
}

// NewNode creates a new Node
func NewNode(id int, u *url.URL, html bool, title string) Node {
	return Node{
		id:   id,
		url:  u,
		html: html,
	}
}

// ID returns the node id
func (n Node) ID() int {
	return n.id
}

// ID returns the node id
func (n Node) DOTID() string {
	return n.url.String()
}

// Edge represents an edge in the directed graph
type Edge struct {
	F, T graph.Node
	W    float64
}

// From returns the source node of this edge
func (e Edge) From() graph.Node {
	return e.F
}

// To returns the destination node of this edge
func (e Edge) To() graph.Node {
	return e.T
}

// Weight returns the weight of this edge
func (e Edge) Weight() float64 {
	return e.W
}
