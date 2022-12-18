package main

import (
	"fmt"
	"net/http"
)

// 案例 函数可以被显式转型
func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome, Gopher!\n")
}
