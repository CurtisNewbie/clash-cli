package clashcli

import "github.com/curtisnewbie/miso/miso"

func GetConfigs(rail miso.Rail, host string) (string, error) {
	return miso.NewTClient(rail, host+"/configs").
		Get().
		Str()
}

func GetProxies(rail miso.Rail, host string) (string, error) {
	return miso.NewTClient(rail, host+"/proxies").
		Get().
		Str()
}

func GetProxiesNamed(rail miso.Rail, host string, name string) (string, error) {
	return miso.NewTClient(rail, host+"/proxies/"+name).
		Get().
		Str()
}

func SelectProxy(rail miso.Rail, host string, proxyGroup string, name string) (string, error) {
	type SelectProxyPayload struct{ Name string }
	dat := SelectProxyPayload{Name: name}
	return miso.NewTClient(rail, host+"/proxies/"+proxyGroup).
		PutJson(dat).
		Str()
}

func GetProxyDelay(rail miso.Rail, host string, name string) (string, error) {
	return miso.NewTClient(rail, host+"/proxies/"+name+"/delay").
		AddQueryParams("timeout", "1000").
		AddQueryParams("url", "https://www.google.com/").
		Get().
		Str()
}
