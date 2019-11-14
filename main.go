package main

import (
	"github.com/google/tcpproxy"
	"github.com/voxelbrain/goptions"
)

func main() {

	options := struct {
		Listen  string `goptions:"-l, --listen, description='will listen port'"`
		ProxyTo string `goptions:"-p, --proxy, description='will proxy to, eg: 192.168.1.23:8888'"`
	}{}

	goptions.ParseAndFail(&options)

	var p tcpproxy.Proxy
	p.AddRoute(options.Listen, tcpproxy.To(options.ProxyTo))

	p.Run()

}
