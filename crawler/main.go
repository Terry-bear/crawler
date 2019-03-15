package main

import (
	"github.com/t496971418/crawler/crawler/config"
	"github.com/t496971418/crawler/crawler/engine"
	"github.com/t496971418/crawler/crawler/persist"
	"github.com/t496971418/crawler/crawler/scheduler"
	"github.com/t496971418/crawler/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
