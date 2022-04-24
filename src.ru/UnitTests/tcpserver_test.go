package tcpserv_test

import (
	"GOLang/src.ru/tcpserv"
	"testing"
)

func TestReadConfig(t *testing.T) {
	var serv tcpserv.TCPServer
	tcpserv.ReadXMLConfig("server.xml", &serv)
	if serv.ServerIP != "127.0.0.1" {
		t.Error("Error Read XML Config : ServerIP != 127.0.0.1")
	}
	if serv.ServerPort != "7777" {
		t.Error("Error Read XML Config : ServerPort != 7777")
	}
}

func TestIntiTCPServer(t *testing.T) {
	var serv tcpserv.TCPServer

	err := serv.ServerStart()
	if err != nil {
		t.Error("Error ServerStart", err)
	}
}
