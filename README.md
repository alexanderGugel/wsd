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
Usage of ./wsd:
  -help
      Display help information about wsd
  -insecureSkipVerify
      Skip TLS certificate verification
  -origin string
      origin of WebSocket client (default "http://localhost/")
  -protocol string
      WebSocket subprotocol
  -url string
      WebSocket server address to connect to (default "ws://localhost:1337/ws")
  -version
      Display version number
```

### Piping

One can piping multiples messages by breaking them with new lines  
E.g.
```
echo 'a
b' \
| wsd -url="wss://echo.websocket.org"
```
results
```
> a
< a
> b
< b
```


## Why?

Debugging WebSocket servers should be as simple as firing up `cURL`. No need
for dozens of flags, just type `wsd -url=ws://localhost:1337/ws` and you're
connected.

## License

 MIT
