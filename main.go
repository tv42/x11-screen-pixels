// Command x11-screen-pixels outputs the screen size of the current
// X11 session as WIDTHxHEIGHT.
package main

import (
	"fmt"
	"log"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

func main() {
	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	fmt.Printf("%dx%d\n", screen.WidthInPixels, screen.HeightInPixels)
}
