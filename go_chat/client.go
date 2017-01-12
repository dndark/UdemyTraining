package main

import (
	"github.com/gorilla/websocket"
	"time"
)

type client struct {
	// socket is the web socket for this client
	socket *websocket.Conn
	// send is a channel on which messages are sent.
	send chan *message
	// room *room
	room *room
	// userata holds information about the user
	userData map[string]interface{}
}

func (c *client) read() {
	defer c.socket.Close()
	for {
		// _, msg, err := c.socket.ReadMessage()
		// if err != nil {
		// 	return
		// }
		// c.room.forward <- msg
		var msg *message
		err := c.socket.ReadJSON(&msg)
		if err != nil {
			return
		}
		msg.When = time.Now()
		msg.Name = c.userData["name"].(string)
		c.room.forward <- msg

	}
}
func (c *client) write() {
	defer c.socket.Close()
	for msg := range c.send {
		err := c.socket.WriteJSON(msg)
		if err != nil {
			return
		}
	}
}
