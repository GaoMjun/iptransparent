module github.com/GaoMjun/iptransparent

go 1.19

replace github.com/GaoMjun/ladder => ../ladder

replace github.com/GaoMjun/tcpip => ../tcpip

require (
	github.com/GaoMjun/ladder v0.0.0-00010101000000-000000000000
	github.com/GaoMjun/tcpip v0.0.0-20230216092336-caa5277ffa13
)

require (
	github.com/GaoMjun/goutils v0.0.0-20230216093315-00bfdc8f95e5 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
)
