package main

import (
	"fmt"

	"log"

	"flag"

	"github.com/t496971418/crawler/crawler/fetcher"
	"github.com/t496971418/crawler/crawler_distributed/rpcsupport"
	"github.com/t496971418/crawler/crawler_distributed/worker"
)

var port = flag.Int("port", 0,
	"the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(
		fmt.Sprintf(":%d", *port),
		worker.CrawlService{}))
}
