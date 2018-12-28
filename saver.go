package main

import (
	"fmt"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
	// debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
	// timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
	// ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
	// count   = kingpin.Arg("count", "Number of packets to send").Int()
	file = kingpin.Flag("text", "Ascii art file to use as screen saver").Required().Short('t').ExistingFile()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Println(file)
}
