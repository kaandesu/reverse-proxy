package server

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"reverse-proxy/internal/configs"
)

func Run() error {
	config, err := configs.NewConfig()
	if err != nil {
		return fmt.Errorf("could not load config:%s", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/find", ping)

	for _, res := range config.Resources {
		url, _ := url.Parse(res.Destination_URL)
		proxy := NewProxy(url)
		mux.HandleFunc(res.Endpoint, ProxyRequestHandler(proxy, url, res.Endpoint))
	}

	if err := http.ListenAndServe(net.JoinHostPort(config.Server.Host, config.Server.Listen_Port), mux); err != nil {
		return fmt.Errorf("could not start the server:%s", err)
	}

	return nil
}
