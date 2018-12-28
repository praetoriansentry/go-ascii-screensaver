package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/nsf/termbox-go"
)

var (
	// debug   = kingpin.Flag("debug", "Enable debug mode.").Bool()
	// timeout = kingpin.Flag("timeout", "Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
	// ip      = kingpin.Arg("ip", "IP address to ping.").Required().IP()
	// count   = kingpin.Arg("count", "Number of packets to send").Int()
	fileName = kingpin.Flag("text", "Ascii art file to use as screen saver").Required().Short('t').ExistingFile()
)

type LogoFile struct {
	File *os.File
	Data [][]rune
	MaxX int
	MaxY int
}

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	logoFile := readFile(*fileName)
	parsedLogo := parseLogo(logoFile)
	fmt.Println(parsedLogo)
	beginLoop(parsedLogo)
}

func readFile(fileName string) *os.File {
	file, err := os.Open(fileName) // For read access.
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func parseLogo(logoFile *os.File) LogoFile {
	l := LogoFile{}
	l.File = logoFile
	l.Data = make([][]rune, 0)
	l.MaxX = 0
	l.MaxY = 0

	scanner := bufio.NewScanner(logoFile)
	for scanner.Scan() {
		t := scanner.Text()
		strlen := len(t)
		if strlen > l.MaxX {
			l.MaxX = strlen
		}
		l.Data = append(l.Data, []rune(t))
	}

	l.MaxY = len(l.Data)
	return l
}

func beginLoop(l LogoFile) {
	err := termbox.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()
	w, h := termbox.Size()

	for {
		drawAt(w, h, l)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				os.Exit(1)
			}
		case termbox.EventResize:
			w, h = termbox.Size()
		}
		time.Sleep(time.Second)
	}
}

func drawAt(xOffset int, yOffset int, l LogoFile) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	y := 0
	for _, lines := range l.Data {
		strlen := len(lines)
		x := 0
		for i := 0; i < strlen; i += 1 {
			termbox.SetCell(x, y, lines[i], termbox.ColorWhite, termbox.ColorDefault)
			x += 1
		}
		y += 1
	}
	termbox.Flush()
}
