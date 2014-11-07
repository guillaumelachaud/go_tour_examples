package main

import (
	"fmt"
	"net/http"
	"strings"
	"utils"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		elem, ok := m[v]
		if ok {
			m[v] = elem + 1
		} else {
			m[v] = 1
		}
	}
	return m
}

func main() {
}

func init() {
	http.HandleFunc("/wc", wc)
}

func wc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, utils.Test(WordCount))
}
