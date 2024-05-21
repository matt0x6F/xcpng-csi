package xolib

import (
	"net/url"

	"github.com/gorilla/websocket"
)

func (xo *xolib) getURL() string {
	scheme := "ws"
	if xo.Config.SSL {
		scheme = "wss"
	}

	host := url.URL{
		Scheme: scheme,
		Host:   xo.Config.Host,
		Path:   "/api/",
	}

	return host.String()
}

func (xo *xolib) getWS() (*websocket.Conn, error) {
	host := xo.getURL()
	ws, _, err := websocket.DefaultDialer.Dial(host, nil)
	return ws, err
}
