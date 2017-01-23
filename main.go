package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/fatih/color"
	"golang.org/x/net/websocket"
)

// Current version number
const Version = "0.0.0"

var (
	origin         string
	url            string
	protocol       string
	displayHelp    bool
	displayVersion bool
	isPipe         bool
	red            = color.New(color.FgRed).SprintFunc()
	magenta        = color.New(color.FgMagenta).SprintFunc()
	green          = color.New(color.FgGreen).SprintFunc()
	yellow         = color.New(color.FgYellow).SprintFunc()
	cyan           = color.New(color.FgCyan).SprintFunc()
	wg             sync.WaitGroup
	wgReceive      sync.WaitGroup
)

func init() {
	flag.StringVar(&origin, "origin", "http://localhost/", "origin of WebSocket client")
	flag.StringVar(&url, "url", "ws://localhost:1337/ws", "WebSocket server address to connect to")
	flag.StringVar(&protocol, "protocol", "", "WebSocket subprotocol")
	flag.BoolVar(&displayHelp, "help", false, "Display help information about wsd")
	flag.BoolVar(&displayVersion, "version", false, "Display version number")
}

func inLoop(ws *websocket.Conn, errors chan<- error, in chan<- []byte) {
	var msg = make([]byte, 512)

	for {
		var n int
		var err error

		n, err = ws.Read(msg)

		if err != nil {
			errors <- err
			continue
		}

		in <- msg[:n]
	}
}

func printErrors(errors <-chan error) {
	for err := range errors {
		if err == io.EOF {
			fmt.Printf("\r✝ %v - connection closed by remote\n", magenta(err))
			os.Exit(0)
		} else {
			fmt.Printf("\rerr %v\n> ", red(err))
			doneReceive()
		}
	}
}

func printReceivedMessages(in <-chan []byte) {
	for msg := range in {
		fmt.Printf("\r< %s\n> ", cyan(string(msg)))
		doneReceive()
	}
}

func outLoop(ws *websocket.Conn, out <-chan []byte, errors chan<- error) {
	for msg := range out {
		_, err := ws.Write(msg)
		if err != nil {
			errors <- err
		}
	}
}

func doneReceive() {
	if isPipe {
		wgReceive.Done()
	}
}

func main() {
	flag.Parse()

	if displayVersion {
		fmt.Fprintf(os.Stdout, "%s version %s\n", os.Args[0], Version)
		os.Exit(0)
	}

	if displayHelp {
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(0)
	}

	ws, err := websocket.Dial(url, protocol, origin)

	if protocol != "" {
		fmt.Printf("connecting to %s via %s from %s...\n", yellow(url), yellow(protocol), yellow(origin))
	} else {
		fmt.Printf("connecting to %s from %s...\n", yellow(url), yellow(origin))
	}

	defer ws.Close()

	if err != nil {
		panic(err)
	}

	fmt.Printf("successfully connected to %s\n\n", green(url))

	wg.Add(3)

	errors := make(chan error)
	in := make(chan []byte)
	out := make(chan []byte)

	defer close(errors)
	defer close(out)
	defer close(in)

	go inLoop(ws, errors, in)
	go printReceivedMessages(in)
	go printErrors(errors)
	go outLoop(ws, out, errors)

	fi, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	isPipe = fi.Mode()&os.ModeNamedPipe != 0

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan() {
		text := scanner.Text()
		if isPipe {
			fmt.Println(green(text))
			wgReceive.Add(1)
		}
		out <- []byte(text)
		if isPipe {
			wgReceive.Wait()
			continue
		}
		fmt.Print("> ")
	}

	wg.Wait()
}
