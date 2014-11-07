package main

import (
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/fib", fib)

}

func wc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, utils.Test(WordCount))
}

func fib(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	f := utils.Fibonacci()
	param := r.URL.Query()["level"][0]
	paramInt, err := strconv.ParseInt(param, 10, 0)

	if err == nil {
		for i := 0; i < int(paramInt); i++ {
			fmt.Fprint(w, f())
			fmt.Fprint(w, "\n")
		}
	} else {
		fmt.Fprint(w, "Invalid parameter !")
	}

}
