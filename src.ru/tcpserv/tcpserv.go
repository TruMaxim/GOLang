package tcpserv

import (
	"GOLang/src.ru/tcpserv/netPackets"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

type TCPServer struct {
	ServerIP   string
	ServerPort string
	ServerName string
	TestServer bool
}

func (serv TCPServer) ServerStart() error {
	fmt.Println("Start", serv.ServerName)

	ReadXMLConfig("server.xml", &serv)
	SettingsLoad()
	if ServerSettings.DebugServer {
		fmt.Println("Debug Mode is ON")
	}

	if serv.ServerIP == "" {
		fmt.Println("ERROR : ServerIP Empty")
		return errors.New("ERROR : ServerIP Empty")
	}

	if serv.ServerPort == "" {
		fmt.Println("ERROR : ServerPort Empty")
		return errors.New("ERROR : ServerPort Empty")
	}

	l, err := net.Listen("tcp", serv.ServerIP+":"+serv.ServerPort)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return err
	}
	defer l.Close()

	fmt.Println("Listening on " + serv.ServerIP + ":" + serv.ServerPort)

	for !serv.TestServer {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return err
		}

		go newClientConnect(conn)
	}

	return nil
}

func newClientConnect(conn net.Conn) {
	var length uint32 = 0
	buf := make([]byte, 65535)
	if ServerSettings.DebugServer {
		fmt.Println("New Client Connection", conn.RemoteAddr())
	}
	dataBuffer := new(bytes.Buffer)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("Error reading:", err.Error())
				break
			}
		}
		m, err := dataBuffer.Write(buf[0:n])
		if err != nil {
			println("ERROR Write To Buffer", m, err.Error())
		}

		// Check Packet Length
		if length == 0 && dataBuffer.Len() >= 4 {
			binary.Read(dataBuffer, binary.LittleEndian, &length)
		}

		if length > 0 && dataBuffer.Len() >= int(length) {
			var packID uint16
			binary.Read(dataBuffer, binary.LittleEndian, &packID)
			// Create New Packet
			pack := netPackets.NewPacketReader(dataBuffer.Next(int(length)), packID, length)

			if ServerSettings.DebugServer {
				fmt.Printf("PackID=%x PackLength=%d ", packID, length)
				fmt.Println("Packet Data :", pack.Data.Bytes())
			}
			length = 0

			// Send Test Packet
			sendpack := netPackets.NewPacketWriter(0x0010)
			// Fill Data
			sendpack.WriteInt32(7777)
			sendpack.WriteFloat32(3.14)
			sendpack.WriteString("1234567890 asdsEWDS авкнАВПВР %#$@$%+}")
			ba := []byte{12, 0, 37, 245, 127}
			sendpack.WriteByteArray(ba)
			sendpack.Send(conn)
		}
	}
}
