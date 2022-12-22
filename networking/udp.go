package networking

import (
	"ethertemp/db"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

var udpServer *net.PacketConn

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
			datas := strings.Split(string(effectiveBuf), "\n")
			temp, err := strconv.ParseFloat(datas[0], 32)
			if err != nil {
				fmt.Println("error converting", datas[0], "to temperature")
				continue
			}

			rh, err := strconv.ParseFloat(datas[1], 32)
			if err != nil {
				fmt.Println("error converting", datas[1], "to humidity")
				continue
			}

			db.AddData(float32(temp), float32(rh))
		}
	}()
}

func Close() {
	_ = (*udpServer).Close()
}
