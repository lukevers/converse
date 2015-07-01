package main

import (
	"database/sql"
	"fmt"
	"github.com/thoj/go-ircevent"
	"log"
	"time"
)

type Client struct {
	Id        int            `db:"id"`
	Nickname  string         `db:"nickname"`
	Realname  string         `db:"realname"`
	Password  sql.NullString `db:"password"`
	Host      string         `db:"host"`
	Port      int            `db:"port"`
	Ssl       bool           `db:"ssl"`
	CreatedAt time.Time      `db:"created_at"`
	conn      *irc.Connection
}

// The exported func Connect is called when connecting a new IRC Client to a
// server. This function should take care of disconnects/reconnects so it's not
// required to worry about that.
func (c *Client) Connect() {
	// Set info from our struct
	c.conn = irc.IRC(c.Nickname, c.Realname)
	c.conn.UseTLS = c.Ssl
	c.conn.Password = c.Password.String

	// Set other options if not running in production mode
	c.conn.Debug = !*production
	c.conn.VerboseCallbackHandler = !*production

	// Figure out address and connect
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)
	err := c.conn.Connect(addr)
	if err != nil {
		log.Println(err)
	}

	// Add event handlers
	c.conn.AddCallback("001", c.HandleWelcome)

	// Stay connected
	c.conn.Loop()
}

// The exported func HandleWelcome is an event handler for irc code "001" which
// handles the welcome event when connecting to an IRC server.
func (c *Client) HandleWelcome(e *irc.Event) {
	// TODO
}
