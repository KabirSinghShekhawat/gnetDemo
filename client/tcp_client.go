package main

import (
	"fmt"
	"net"
	"os"
	"sync"
)

func sendRequest(host, port string, reqNumber int, wg *sync.WaitGroup) {
	defer wg.Done()

	tcpServer, err := net.ResolveTCPAddr("tcp", host+":"+port)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(fmt.Sprintf("%d: Hello, World!", reqNumber)))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	defer conn.Close()
}

func main() {
	const (
		HOST = "localhost"
		PORT = "9000"
	)

	numberOfRequests := 10000

	var wg sync.WaitGroup

	for i := 0; i < numberOfRequests; i++ {
		wg.Add(1)
		reqNumber := i
		go sendRequest(HOST, PORT, reqNumber, &wg)
	}

	wg.Wait()
}
