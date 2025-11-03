package models

import "github.com/gofiber/contrib/websocket"

type ClientRegistration struct {
	ID   uint
	Conn *websocket.Conn
}
