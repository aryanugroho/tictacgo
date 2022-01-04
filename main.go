package main

import (
	"github.com/aryanugroho/tictacgo/transport"
)

func main() {
	app := transport.NewCliTransport()
	app.Run()
}
