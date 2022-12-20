package main

import (
	"bufio"
	"ethertemp/db"
	"ethertemp/networking"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalln("could not load environment file")
	}

	err = db.InitializeMongo()
	if err != nil {
		log.Fatalln("error connecting to db:", err.Error())
		return
	}

	defer db.Close()

	err = networking.InitializeUdpSocket()
	if err != nil {
		log.Fatalln("error listening for udp:", err.Error())
		return
	}

	networking.StartListener()

	fmt.Println("udp server started")

	defer networking.Close()

	fmt.Println("type q to turn off the server")
	consoleReader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := consoleReader.ReadLine()
		if err != nil {
			fmt.Println("error reading stdin line, error:", err.Error())
		}

		if string(line) == "q" {
			return
		}
	}
}
