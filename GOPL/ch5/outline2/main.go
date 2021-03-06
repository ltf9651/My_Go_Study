package main

import "My_Go_Study/GOPL/ch5/html"

func forEachNode(n *html.Node, pre, post func(n, *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
