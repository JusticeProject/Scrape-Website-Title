package main

import "strings"

func ExtractTitle(html string) string {
	hasTitle := strings.Contains(html, "<title>") && strings.Contains(html, "</title>")
	if !hasTitle {
		return ""
	}

	html_split := strings.Split(html, "<title>")
	title := strings.Split(html_split[1], "</title>")[0]
	title = strings.ReplaceAll(title, "\n", "")
	title = strings.ReplaceAll(title, "\r", "")
	title = strings.ReplaceAll(title, "\t", "")
	title = strings.Trim(title, " ")
	return title
}
