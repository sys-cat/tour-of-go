package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, f := range strings.Fields(s) { // strings.Fields は英語のみに有効。新語や日本語などには対応していないので現時点では英語のみ対応
		m[f]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
