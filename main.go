// Command x11-screen-pixels outputs the screen size of the current
// X11 session as WIDTHxHEIGHT.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", prog)
	fmt.Fprintf(os.Stderr, "  %s\n", prog)
	fmt.Fprintf(os.Stderr, "  DISPLAY=:1 %s\n", prog)
	fmt.Fprintf(os.Stderr, "(the command takes no options)\n")
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")

	flag.Usage = usage
	flag.Parse()
	if flag.NArg() != 0 {
		flag.Usage()
		os.Exit(2)
	}

	X, err := xgb.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	setup := xproto.Setup(X)
	screen := setup.DefaultScreen(X)
	fmt.Printf("%dx%d\n", screen.WidthInPixels, screen.HeightInPixels)
}
