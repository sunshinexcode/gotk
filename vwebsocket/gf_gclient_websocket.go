package vwebsocket

import "github.com/gogf/gf/v2/net/gclient"

type (
	WebSocketClient = gclient.WebSocketClient
)

func NewClient() *WebSocketClient {
	return gclient.NewWebSocket()
}
