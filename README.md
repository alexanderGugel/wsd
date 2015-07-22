# wsd

> = **W**eb**S**ocket **D**ebugger

[![Build Status](https://travis-ci.org/alexanderGugel/wsd.svg?branch=master)](https://travis-ci.org/alexanderGugel/wsd)

![Terminal Demo](https://cdn.rawgit.com/alexanderGugel/wsd/demo/demo.gif)

Simple command line utility for debugging WebSocket servers.

## Installation

Via `go-get`:

```
$ go get github.com/alexanderGugel/wsd
```

## Usage

Command-line usage:

```
  -help=false: Display help information about wsd
  -origin="http://localhost/": origin of WebSocket client
  -protocol="": WebSocket subprotocol
  -url="ws://localhost:1337/ws": WebSocket server address to connect to
  -version=false: Display version number
```

## Why?

Debugging WebSocket servers should be as simple as firing up `cURL`. No need
for dozens of flags, just type `wsd -url=ws://localhost:1337/ws` and you're
connected.

## License

 MIT
