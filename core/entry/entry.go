/**
* @Author: shaochuyu
* @Date: 5/7/2022 11:30
 */
package entry

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/thoas/go-funk"
	"github.com/urfave/cli/v2"
	"go.uber.org/automaxprocs/maxprocs"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strings"
	"wscan/core/collector"
	"wscan/core/collector/basiccrawler"
	"wscan/core/crawler"
	"wscan/core/ctrl"
	"wscan/core/http"
	"wscan/core/plugins"
	"wscan/core/plugins/base"
	"wscan/core/utils"
	"wscan/core/utils/checker"
	"wscan/core/utils/log"
	"wscan/core/utils/printer"
)

func NewApp(c *cli.Context) {
	cfg, err := LoadOrGenConfig(c)
	if err != nil {
		log.Fatal(err)
	}
	if c.Bool("dump-config") == true {
		log.Info("Dumping example config to ./config.yaml.example")
	}
	maxprocs.Set()
	var col collector.Fitter
	targets := []string{}
	if c.String("basic-crawler") != "" {
		col = basiccrawler.NewBasicCrawlerCollector(cfg.HTTP, &crawler.Config{
			Browser:                false,
			RestrictionsOnRequests: crawler.RestrictionsOnRequests{MaxConcurrent: 5, MaxDepth: 0},
			Restrictions:           cfg.Filter,
		})
		targets = append(targets, c.String("basic-crawler"))
	} else if c.String("browser-crawler") != "" {
		col = basiccrawler.NewBasicCrawlerCollector(cfg.HTTP, &crawler.Config{
			ExecPath:        cfg.BrowserConfig.ExecPath,
			DisableHeadless: cfg.BrowserConfig.DisableHeadless,
			Browser:         true,
			RestrictionsOnRequests: crawler.RestrictionsOnRequests{
				MaxConcurrent:  cfg.BrowserConfig.MaxPageConcurrent,
				MaxDepth:       cfg.BrowserConfig.MaxDepth,
				MaxCountOfURLs: cfg.BrowserConfig.MaxPageVisit,
			},
			Restrictions: cfg.Filter,
		})
		targets = append(targets, c.String("browser-crawler"))
	} else if c.String("url") != "" {
		col = collector.NewFromURLListReader(ioutil.NopCloser(strings.NewReader(c.String("data"))), cfg.HTTP)
		targets = append(targets, c.String("url"))
	} else if c.String("raw-request") != "" {

	} else if c.String("url-file") != "" {
		col = collector.NewFromURLListReader(ioutil.NopCloser(strings.NewReader(c.String("data"))), cfg.HTTP)
		urlFile, err := os.Open(c.String("url-file"))
		if err != nil {
			log.Fatal(err)
		}
		defer urlFile.Close()
		scanner := bufio.NewScanner(urlFile)
		for scanner.Scan() {
			url := strings.TrimSpace(scanner.Text())
			if url != "" && funk.ContainsString(targets, url) == false {
				targets = append(targets, url)
			}
		}
	} else if c.String("listen") != "" {
		log.Println("listen=", c.String("listen"))
		cfg.Mitm.Listen = c.String("listen")
		col = collector.NewMitmProxy(&cfg.Mitm, cfg.HTTP)
	} else {
		log.Fatal("Warning: you should use --html-output, --webhook-output or --json-output to persist your scan result\n`url`, `listen`, `raw-request`, `url-file`, `basic-crawler` and `browser-crawler` must use one, try `--help` to see details")
	}
	taskChan, err := col.FitOut(context.Background(), targets)
	if err != nil {
		log.Fatal(err)
	}
	multiPrinter := printer.NewMultiPrinter()
	printers := getPrinters(c)
	multiPrinter.AddPrinters(printers)
	dispatcher := ctrl.NewDispatcher(&cfg.Config, multiPrinter)
	dispatcher.Init(false)
	dispatcher.Run(taskChan)
	dispatcher.Release()

	defer multiPrinter.Close()
}

func LoadOrGenConfig(c *cli.Context) (*CliEntryConfig, error) {
	configPath := c.String("config")
	if configPath == "" {
		configPath = "./config.yaml"
	}
	cfg := &CliEntryConfig{}
	if !utils.FileExists(configPath) {
		log.Info("Generate default configurations to config.yaml")
		cfg = NewExampleConfig()
		cfgData, err := yaml.Marshal(cfg)
		if err != nil {
			return nil, err
		}
		// Generate default config
		if err = ioutil.WriteFile("./config.yaml", cfgData, 0644); err != nil {
			log.Error("can't write default config to config.yaml, please check permission.")
			return nil, err
		}
	} else {
		cfgData, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Error(err)
			return nil, err
		}

		err = yaml.Unmarshal(cfgData, &cfg)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		cfg.Config.Plugins = make(map[string]base.PluginConfigInterface)
		for _, p := range plugins.All() {
			pluginConfig := p.DefaultConfig()
			jsonData, err := json.Marshal(cfg.Plugins[pluginConfig.BaseConfig().Name])
			if err != nil {
				log.Error(err)
				continue
			}
			if err = json.Unmarshal(jsonData, pluginConfig); err != nil {
				log.Error(err)
				continue
			}
			cfg.Config.Plugins[pluginConfig.BaseConfig().Name] = pluginConfig
		}
	}

	if c.String("plugins") != "" {
		enabledPlugins := strings.Split(c.String("plugins"), ",")
		for name, pc := range cfg.Config.Plugins {
			if funk.ContainsString(enabledPlugins, name) == true {
				pc.BaseConfig().Enabled = true
			} else {
				pc.BaseConfig().Enabled = false
			}
		}
	}

	cfg.Filter = &checker.RequestCheckerConfig{
		URLCheckerConfig: checker.URLCheckerConfig{
			SchemeAllowed:        []string{},
			SchemeDisallowed:     []string{},
			HostnameAllowed:      []string{},
			HostnameDisallowed:   []string{},
			TCPPortAllowed:       []string{},
			TCPPortDisallowed:    []string{},
			PathAllowed:          []string{},
			PathDisallowed:       []string{},
			PathSuffixAllowed:    []string{},
			PathSuffixDisallowed: []string{},
			QueryKeyAllowed:      []string{},
			QueryKeyDisallowed:   []string{},
			QueryRawAllowed:      []string{},
			QueryRawDisallowed:   []string{},
			FragmentAllowed:      []string{},
			FragmentDisallowed:   []string{},
			URLRegexAllowed:      []string{},
			URLRegexDisallowed:   []string{},
			URLGlobAllowed:       []string{},
			URLGlobDisallowed:    []string{},
		},
		MethodAllowed:     []string{},
		MethodDisallowed:  []string{},
		PostKeyAllowed:    []string{},
		PostKeyDisallowed: []string{},
	}

	logLevel := c.String("log-level")
	logCfg := log.NewDefaultConfig()
	if logLevel != "" {
		logCfg.Level = logLevel
	}
	log.SetConfig(&logCfg)
	clientOptions := http.ClientOptions{}
	clientOptions.WroteBack()
	cfg.Subdomain.WroteBack()
	return cfg, nil

}

func init() {

}
