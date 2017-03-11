package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/uudashr/fibweb"
	"github.com/uudashr/fibweb/httpfib"
)

var usageMsg = `
Usage: fibweb [options]
     --fibgo-addr   (required) Fibgo address ex: http://fibgo:8080
 -p, --port         Listen port
 -h, --help         Show this message
`

func usage() {
	fmt.Println(usageMsg)
}

func main() {
	var port int
	var fibgoAddr string
	var showHelp bool

	flag.IntVar(&port, "port", 8080, "Listen port")
	flag.IntVar(&port, "p", 8080, "Listen port")
	flag.StringVar(&fibgoAddr, "fibgo-addr", "", "Fibgo service address")
	flag.BoolVar(&showHelp, "help", false, "Show this message")
	flag.BoolVar(&showHelp, "h", false, "Show this message")

	flag.Usage = usage
	flag.Parse()

	if showHelp || fibgoAddr == "" {
		flag.Usage()
		return
	}

	fibService := httpfib.NewFibonacciService(fibgoAddr)
	handler := fibweb.NewHTTPHandler(fibService)
	log.Println("Listen on port", port, "...")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler)
	log.Println("Stopped err:", err)
}
