package main

import (
    "github.com/gorilla/websocket"
)

type client struct {
    // socket is the web socket for tis client
    socket *websocket.Conn
    // send is a channel on which messages are sent.
    send chan []byte
    // room *room
    room *room
}

type room struct {
    // forward is a channel that holds incoming messages
    // that should be forwarded to the other clients.
    forward chan []byte
}
