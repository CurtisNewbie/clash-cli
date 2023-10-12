# clash-cli

CLI utils to talk to clash

doc: https://dreamacro.github.io/clash/runtime/external-controller.html


```sh
go run main.go -host http://curtisnewbie.com:9090 -command GetConfigs
go run main.go -host http://curtisnewbie.com:9090 -command GetProxies
go run main.go -host http://curtisnewbie.com:9090 -command GetProxiesNamed -proxy-name Proxy
go run main.go -host http://curtisnewbie.com:9090 -command SelectProxy -proxy-group Proxy -proxy-name "V3-373|香港|x1.5"
go run main.go -host http://curtisnewbie.com:9090 -command GetProxyDelay -proxy-name "V3-373|香港|x1.5"
```