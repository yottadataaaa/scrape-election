package main

import (
	"context"
	"log"

	"scrape-election/config"
	"scrape-election/crawler"
	"scrape-election/loglib"
)

func main() {
	cfg := config.InitConfig()

	logger := log.New(log.Writer(), "[MyApp] ", log.LstdFlags)
	ctx := context.WithValue(context.Background(), loglib.LoggerKey, logger)

	crawler := crawler.Crawler{}
	err := crawler.DoProcess(ctx, cfg)
	if err != nil {
		// TODO: エラーを吐き出す
		panic(err)
	}
}
