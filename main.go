package main

import (
	"flag"
	"github.com/mhthrh/ApiStore/View"
)

func main() {
	addr := flag.String("addr", "localhost:8080", "the TCP address for the server to listen on, in the form 'host:port'")
	View.RunApi(*addr)
}
