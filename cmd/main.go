package main

import "github.com/luchojuarez/call-sorter/cmd/web"

func main() {
	web.NewWebServer().ListenAndServe()
}
