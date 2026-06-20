package main

import "strings"

func cleanInput(text string) []string {
	text = strings.TrimSpace(text)
	text = strings.ToLower(text)
	cut := strings.Fields(text)
	return cut
}