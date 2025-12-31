/*
*
2 * @Author: shaochuyu
3 * @Date: 8/24/25
4
*/
package mcp

import (
	"context"
	"io"
	"strings"
	"wscan/core/collector"
	"wscan/core/collector/basiccrawler"
	"wscan/core/crawler"
	"wscan/core/ctrl"
	"wscan/core/output"
	logger "wscan/core/utils/log"
	"wscan/core/utils/printer"
)

func UrlScan(task *Task, mode string) error {
	var col collector.Fitter
	rConfig := crawler.RequestConfig{}
	for k, v := range globalConfig.HTTP.HEADER_NO_USE {
		rConfig.Headers = append(rConfig.Headers, crawler.Header{
			k, v,
		})
	}
	if mode == "basic" {
		col = basiccrawler.NewBasicCrawlerCollector(globalConfig.HTTP, &crawler.Config{
			Proxy:                  globalConfig.HTTP.Proxy,
			Browser:                false,
			RestrictionsOnRequests: crawler.RestrictionsOnRequests{MaxConcurrent: 5, MaxDepth: 0},
			Restrictions:           globalConfig.BasicCrawler.Restriction,
			RequestConfig:          rConfig,
		})
	} else if mode == "browser" {
		col = basiccrawler.NewBasicCrawlerCollector(globalConfig.HTTP, &crawler.Config{
			Proxy:           globalConfig.HTTP.Proxy,
			ExecPath:        globalConfig.BrowserConfig.ExecPath,
			DisableHeadless: globalConfig.BrowserConfig.DisableHeadless,
			Browser:         true,
			RestrictionsOnRequests: crawler.RestrictionsOnRequests{
				MaxConcurrent:  globalConfig.BrowserConfig.MaxPageConcurrent,
				MaxDepth:       globalConfig.BrowserConfig.MaxDepth,
				MaxCountOfURLs: globalConfig.BrowserConfig.MaxPageVisit,
			},
			Restrictions:  globalConfig.BrowserConfig.Restriction,
			RequestConfig: rConfig,
		})
	} else {
		col = collector.NewFromURLListReader(io.NopCloser(strings.NewReader("")), globalConfig.HTTP)
	}
	logger.Infof("url scan task:%v crawler type %s", task.ScanUrl, mode)
	targets := []string{task.ScanUrl}
	var err error
	taskChan, err := col.FitOut(context.Background(), targets)
	if err != nil {
		logger.Fatal(err)
	}

	multiPrinter := printer.NewMultiPrinter()
	printers := []printer.Printer{}
	printers = append(printers, NewMcpPrinter(task))
	printers = append(printers, output.NewStdoutPrinter())
	multiPrinter.AddPrinters(printers)

	dispatcher := ctrl.NewDispatcher(&globalConfig.Config, multiPrinter, globalReverse)
	dispatcher.Init(false)
	dispatcher.Run(taskChan, false)
	dispatcher.Release()

	defer multiPrinter.Close()

	logger.Infof("task name %s  id %s finished", task.Name, task.ID)
	return nil
}
