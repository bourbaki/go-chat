package main

import (
	"io"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		c, err := l.Accept()
		println(c.RemoteAddr().String())
		if err != nil {
			log.Fatal(err)
		}
		go io.Copy(c, c) // Обработчик для каждого клиента
	}
}