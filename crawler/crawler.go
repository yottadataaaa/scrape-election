package crawler

import (
	"context"

	"scrape-election/config"
	"scrape-election/loglib"
)

type Crawler struct {
}

func (c Crawler) DoProcess(
	ctx context.Context,
	cfg *config.ProcessConfig,
) error {
	logger, err := loglib.UseLogger(ctx)
	if err != nil {
		return err
	}

	logger.Println("process started ...")

	return nil
}

func (c Crawler) GetCrawlerProcessor() {

}
