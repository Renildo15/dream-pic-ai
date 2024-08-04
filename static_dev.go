//+build dev
//go:build dev
// +build dev

package main

import (
	"net/http"
	"os"
	"fmt"
)


func public() http.Handler {
	fmt.Println("building static files for development")
	return http.StripPrefix("/public/", http.FileServerFS(os.DirFS("public")))
}