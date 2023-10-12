package main

import (
	"flag"

	"github.com/curtisnewbie/clash-cli/clashcli"
	"github.com/curtisnewbie/miso/miso"
)

const (
	CmdGetConfigs      = "GetConfigs"
	CmdGetProxies      = "GetProxies"
	CmdGetProxiesNamed = "GetProxiesNamed"
	CmdSelectProxy     = "SelectProxy"
	CmdGetProxyDelay   = "GetProxyDelay"
	CmdGetLogs         = "GetLogs"
)

func main() {
	Host := flag.String("host", "http://localhost:9090", "clash host address")
	Command := flag.String("command", CmdGetProxies, "clash command")
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
		r, err := clashcli.GetProxyDelay(rail, *Host, *Proxy)
		if err != nil {
			panic(err)
		}
		rail.Infof("Proxy delay %v: %v", *Proxy, r)
	}
}
