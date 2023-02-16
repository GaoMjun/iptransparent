module github.com/GaoMjun/iptransparent

go 1.19

replace tcpip => github.com/GaoMjun/tcpip v0.0.0-20191023054252-0c623063cd59

replace ladder => github.com/GaoMjun/ladder v0.0.0-20190418104442-14c247b1205e

require (
	ladder v0.0.0-00010101000000-000000000000
	tcpip v0.0.0-00010101000000-000000000000
)

require (
	github.com/GaoMjun/goutils v0.0.0-20230216091549-748a48fb2458 // indirect
	github.com/armon/go-socks5 v0.0.0-20160902184237-e75332964ef5 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	golang.org/x/net v0.7.0 // indirect
)
