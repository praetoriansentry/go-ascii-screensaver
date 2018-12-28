package main

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/nsf/termbox-go"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

var (
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
	rand.Seed(time.Now().Unix())

	go func() {
		for {
			switch ev := termbox.PollEvent(); ev.Type {
			case termbox.EventInterrupt:
				os.Exit(1)
			case termbox.EventKey:
				if ev.Key == termbox.KeyEsc {
					os.Exit(1)
				}
			}
		}
	}()

	for {
		w, h := termbox.Size()
		wRange := w - l.MaxX
		hRange := h - l.MaxY
		if wRange < 1 || hRange < 1 {
			log.Fatal("Screen size is too small for graphic")
		}
		drawAt(rand.Intn(wRange), rand.Intn(hRange), l)
		time.Sleep(time.Second)
		termbox.Sync()
	}

}

func drawAt(xOffset int, yOffset int, l LogoFile) {
	termbox.Clear(termbox.ColorWhite, termbox.ColorDefault)
	y := yOffset
	for _, lines := range l.Data {
		strlen := len(lines)
		x := xOffset
		for i := 0; i < strlen; i += 1 {
			termbox.SetCell(x, y, lines[i], termbox.ColorWhite, termbox.ColorDefault)
			x += 1
		}
		y += 1
	}
	termbox.Flush()
}
