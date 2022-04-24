package main

import "GOLang/src.ru/tcpserv"

var serv tcpserv.TCPServer

func main() {
	serv.ServerName = "Test Server"
	serv.ServerStart()
}
