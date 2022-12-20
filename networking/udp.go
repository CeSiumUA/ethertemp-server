package networking

import (
	"encoding/binary"
	"ethertemp/db"
	"fmt"
	"math"
	"net"
	"os"
)

var udpServer *net.PacketConn

func toFloat32(bytes []byte) float32 {
	bits := binary.LittleEndian.Uint32(bytes)
	return math.Float32frombits(bits)
}

func InitializeUdpSocket() error {
	socket, err := net.ListenPacket("udp", os.Getenv("UDP_PORT"))
	if err != nil {
		return err
	}
	udpServer = &socket
	return err
}

func StartListener() {
	go func() {
		for {
			buf := make([]byte, 1024)
			n, addr, err := (*udpServer).ReadFrom(buf)
			if err != nil {
				fmt.Println("error reading udp socket:", err)
				continue
			}

			effectiveBuf := buf[:n]
			fmt.Println("got package from:", addr.String())
			fmt.Println("data:", string(effectiveBuf))
			tmpResult := toFloat32(effectiveBuf)
			db.AddTemperature(tmpResult)
		}
	}()
}

func Close() {
	_ = (*udpServer).Close()
}
