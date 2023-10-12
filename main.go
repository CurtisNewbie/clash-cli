package main

import (
	"flag"
	"fmt"
	"sync"

	"github.com/curtisnewbie/clash-cli/clashcli"
	"github.com/curtisnewbie/miso/miso"
)

const (
	CmdGetConfigs      = "GetConfigs"
	CmdGetProxies      = "GetProxies"
	CmdGetProxiesNamed = "GetProxiesNamed"
	CmdSelectProxy     = "SelectProxy"
	CmdGetProxyDelay   = "GetProxyDelay"
)

var (
	Commands = []string{
		CmdGetConfigs,
		CmdGetProxies,
		CmdGetProxiesNamed,
		CmdSelectProxy,
		CmdGetProxyDelay,
	}
)

func main() {
	Host := flag.String("host", "http://curtisnewbie.com:9090", "clash host address")
	Command := flag.String("command", "", fmt.Sprintf("clash command: %v", Commands))
	ProxyGroup := flag.String("proxy-group", "Proxy", "proxy group")
	Proxy := flag.String("proxy-name", "", "proxy name")
	flag.Parse()

	rail := miso.EmptyRail()
	rail.SetLogLevel("debug")

	if *Host == "" {
		panic("Host is required")
	}

	switch *Command {
	case CmdGetConfigs:
		r, err := clashcli.GetConfigs(rail, *Host)
		if err != nil {
			panic(err)
		}
		rail.Infof("Configs: %v", r)
	case CmdGetProxies:
		r, err := clashcli.GetProxies(rail, *Host)
		if err != nil {
			panic(err)
		}
		rail.Infof("Proxies: %v", r)
	case CmdGetProxiesNamed:
		r, err := clashcli.GetProxiesNamed(rail, *Host, *Proxy)
		if err != nil {
			panic(err)
		}
		rail.Infof("Proxies: %v", r)
	case CmdSelectProxy:
		r, err := clashcli.SelectProxy(rail, *Host, *ProxyGroup, *Proxy)
		if err != nil {
			panic(err)
		}
		rail.Infof("Selected proxy %v: %v", *Proxy, r)
	case CmdGetProxyDelay:
		GetDelayAll(rail, *Host, *ProxyGroup)
	}
}

type ProxiesNamed struct {
	All  []string
	Name string
	Now  string
}

func GetDelayAll(rail miso.Rail, host string, name string) {
	r, err := clashcli.GetProxiesNamed(rail, host, name)
	if err != nil {
		panic(err)
	}
	rail.Infof("Proxies: %v", r)

	var proxies ProxiesNamed
	if err := miso.ParseJson([]byte(r), &proxies); err != nil {
		panic(err)
	}
	rail.Infof("ProxiedNamed: %+v", proxies)

	var wg sync.WaitGroup

	for i := range proxies.All {
		wg.Add(1)
		go func(p string) {
			defer wg.Done()
			delay, err := clashcli.GetProxyDelay(rail, host, p)
			if err != nil {
				rail.Errorf("Check proxy delay failed, %v, %v", p, err)
				return
			}
			rail.Infof("%v delay %v", p, delay)
		}(proxies.All[i])
	}
	wg.Wait()
}
