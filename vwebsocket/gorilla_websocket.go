package vwebsocket

import "github.com/gorilla/websocket"

type (
	Conn = websocket.Conn
)

const (
	BinaryMessage = websocket.BinaryMessage
	TextMessage   = websocket.TextMessage
)
